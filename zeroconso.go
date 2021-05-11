package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/hd44780i2c"
)

const (
	DISPLAY_NUM_COLS  = 20
	DISPLAY_NUM_LINES = 4

	SCREEN_SPLASH = 0
	SCREEN_INFO   = iota
	SCREEN_CALIB  = iota

	KEY_UP        = 0
	KEY_DOWN      = iota
	KEY_LEFT      = iota
	KEY_RIGHT     = iota
	KEY_SHORT     = iota
	KEY_LONGPRESS = iota
)

type MetricsData struct {
	AcVoltage1 uint32 // AC Voltage Line 1
	AcCurrent1 uint32 // AC Current Line 1
	AcVoltage2 uint32 // AC Voltage Line 2
	AcCurrent2 uint32 // AC Current Line 2
}

type ConfigData struct {
	AcVoltage1Calib uint16 // Calibration for AC Voltage Line 1
	AcCurrentCalib  uint16 // Calibration for AC Current Line 1
	AcVoltageCalib  uint16 // Calibration for AC Voltage Line 2
	AcCurrent2Calib uint16 // Calibration for AC Current Line 2
}

var (
	lcd     hd44780i2c.Device
	metrics MetricsData
	config  ConfigData
)

//var buffer [DISPLAY_NUM_COLS * DISPLAY_NUM_LINES]byte

//  --------- HW INIT

func hwInit() {

	// I2C and LCD
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})
	lcd = hd44780i2c.New(machine.I2C0, 0x27) // some modules have address 0x3F
	lcd.Configure(hd44780i2c.Config{
		Width:       20, // required
		Height:      4,  // required
		CursorOn:    true,
		CursorBlink: true,
	})
}

// ---------- LCD

func DisplayScreen(screen int, cnf ConfigData, rt MetricsData) {
	lcd.SetCursor(0, 0)
	switch screen {
	case SCREEN_SPLASH:
		lcd.Print([]byte("Zero Conso V1"))
	case SCREEN_CALIB:
		lcd.Print([]byte("Calibration"))
	case SCREEN_INFO:
		lcd.Print([]byte("Info"))
	}
}

// ---------- Key
func handleKeyboard(key byte) {

}

// ---------- MAIN

func main() {

	println("Hello TinyGo")

	hwInit()

	// Blink led
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for i := 0; i < 5; i++ {
		led.Low()
		time.Sleep(time.Millisecond * 100)
		led.High()
		time.Sleep(time.Millisecond * 100)
	}
	println("2")

	lcd.SetCursor(0, 0)
	lcd.Print([]byte("TinyGO"))

	//DisplayScreen(SCREEN_SPLASH, config, metrics)

}
