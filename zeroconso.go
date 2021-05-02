package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/hd44780i2c"
)

func main() {

	println("Hello TinyGo")

	// Blink led
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for i := 0; i < 5; i++ {
		led.Low()
		time.Sleep(time.Millisecond * 100)
		led.High()
		time.Sleep(time.Millisecond * 100)
	}

	// I2C and LCD
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	lcd := hd44780i2c.New(machine.I2C0, 0x27) // some modules have address 0x3F
	lcd.Configure(hd44780i2c.Config{
		Width:       16, // required
		Height:      2,  // required
		CursorOn:    false,
		CursorBlink: false,
	})

	lcd.Print([]byte(" TinyGo\n  LCD Test "))

}
