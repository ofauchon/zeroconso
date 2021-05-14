package main

import (
	"machine"
	"time"

	"./core"
	"./drivers"

	"tinygo.org/x/drivers/hd44780i2c"
)

const ()

var (
	rt *core.RuntimeData
)

//  --------- HW INIT

func hwInit() {

	rt.SerialConsole = &machine.UART0
	// I2C and LCD
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})
	driver := hd44780i2c.New(machine.I2C0, 0x27) // some modules have address 0x3F
	rt.Lcd = &driver
	rt.Lcd.Configure(hd44780i2c.Config{
		Width:       20, // required
		Height:      4,  // required
		CursorOn:    false,
		CursorBlink: false,
	})
}

// ---------- Key
func handleKeyboard(key byte) {
	println("key:", key)

}

func readSerial() {
	for {
		if rt.SerialConsole.Buffered() > 0 {
			data, _ := rt.SerialConsole.ReadByte()
			handleKeyboard(data)
		}
		time.Sleep(10 * time.Millisecond)
	}

}

// ---------- MAIN

func main() {

	rt = &core.RuntimeData{}
	rt.Metrics = &core.MetricsData{}
	println("Hello TinyGo")

	hwInit()
	rt.Lcd.SetCursor(0, 0)
	rt.Lcd.Print([]byte("HELLO"))

	go readSerial()

	// Blink led
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	drivers.AdcInit()

	for i := 0; i < 5; i++ {
		led.Low()
		time.Sleep(time.Millisecond * 100)
		led.High()
		time.Sleep(time.Millisecond * 100)
	}

	display := core.NewDisplayManager(rt)
	display.ScreenInfos(true)

	for {

		// Read U1
		drivers.AdcSetChannel(0)
		vms := drivers.AcAmplitudeMv()
		ima := (30000 * vms) / 1000
		rt.Metrics.AcVoltageMv = ima

		// Read I1
		drivers.AdcSetChannel(1)
		vms = drivers.AcAmplitudeMv()
		ima = (30000 * vms) / 1000
		rt.Metrics.AcCurrent1Ma = ima

		// Read I2
		drivers.AdcSetChannel(2)
		vms = drivers.AcAmplitudeMv()
		ima = (30000 * vms) / 1000
		rt.Metrics.AcCurrent2Ma = ima

		// Read I3
		drivers.AdcSetChannel(3)
		vms = drivers.AcAmplitudeMv()
		ima = (30000 * vms) / 1000
		rt.Metrics.AcCurrent3Ma = ima

		display.ScreenInfos(false)
		time.Sleep(time.Second * 5)
	}

}
