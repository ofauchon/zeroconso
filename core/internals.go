package core

import (
	"machine"

	"tinygo.org/x/drivers/hd44780i2c"
)

const (
	DISPLAY_NUM_COLS  = 20
	DISPLAY_NUM_LINES = 4

	KEY_UP        = 0
	KEY_DOWN      = iota
	KEY_LEFT      = iota
	KEY_RIGHT     = iota
	KEY_SHORT     = iota
	KEY_LONGPRESS = iota
)

type MetricsData struct {
	AcVoltageMv  uint32 // AC Voltage Common
	AcCurrent1Ma uint32 // AC Current Line 1
	AcCurrent2Ma uint32 // AC Current Line 2
	AcCurrent3Ma uint32 // AC Current Line 2
}

type ConfigData struct {
	AcVoltage1Calib uint16 // Calibration for AC Voltage Line 1
	AcCurrentCalib  uint16 // Calibration for AC Current Line 1
	AcVoltageCalib  uint16 // Calibration for AC Voltage Line 2
	AcCurrent2Calib uint16 // Calibration for AC Current Line 2
}

type RuntimeData struct {
	Lcd           *hd44780i2c.Device
	Metrics       *MetricsData
	Config        *ConfigData
	SerialConsole *machine.UART
}
