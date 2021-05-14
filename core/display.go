package core

import "strconv"

const (
	SCREEN_SPLASH = 0
	SCREEN_INFO   = iota
	SCREEN_CONF   = iota
	SCREEN_VERS   = iota

	EVT_START     = 1
	EVT_UP        = iota
	EVT_DOWN      = iota
	EVT_LEFT      = iota
	EVT_RIGHT     = iota
	EVT_BT1_SHORT = iota
	EVT_BT1_LONG  = iota
)

type DisplayManager struct {
	rd            *RuntimeData
	currentScreen int
}

var (
	countRefresh int64
)

func u16String(i uint16) []byte {

	v := ""
	if i < 1000 {
		v += "0"
	}
	if i < 100 {
		v += "0"
	}
	if i < 10 {
		v += "0"
	}
	//println("i:", i)
	v += strconv.Itoa(int(i))
	return []byte(v)
}

func formatMillis(val uint32, unit string) string {
	entire := int(val / 1000)
	decimal := int(val%1000) / 100
	v := ""
	if entire < 10 {
		v += " "
	}
	v += strconv.Itoa(entire) + "." + strconv.Itoa(decimal)
	if unit != "" {
		v += unit
	}
	return v
}

func NewDisplayManager(run *RuntimeData) *DisplayManager {
	ret := &DisplayManager{rd: run}
	return ret
}

func (d *DisplayManager) Event(ev int) {

	/*
		if ev == EVT_BT1_SHORT {
			if d.currentScreen == SCREEN_INFO {
				d.currentScreen == SCREEN_CONF
			}
		}
	*/
}

func (d *DisplayManager) ScreenInfos(fullUpdate bool) {
	d.currentScreen = SCREEN_INFO

	if fullUpdate {
		d.rd.Lcd.ClearDisplay()
		d.rd.Lcd.SetCursor(0, 0)
		d.rd.Lcd.Print([]byte("INFO"))
		d.rd.Lcd.SetCursor(0, 1)
		d.rd.Lcd.Print([]byte("U1:"))
		d.rd.Lcd.SetCursor(10, 1)
		d.rd.Lcd.Print([]byte("I1:"))
		d.rd.Lcd.SetCursor(0, 2)
		d.rd.Lcd.Print([]byte("I2:"))
		d.rd.Lcd.SetCursor(10, 2)
		d.rd.Lcd.Print([]byte("I3:"))
	}

	d.rd.Lcd.SetCursor(18, 0)
	if (countRefresh % 2) == 0 {

		d.rd.Lcd.Print([]byte("o"))
	} else {
		d.rd.Lcd.Print([]byte("O"))
	}
	countRefresh++

	d.rd.Lcd.SetCursor(3, 1)
	d.rd.Lcd.Print([]byte(formatMillis(d.rd.Metrics.AcVoltageMv, "V")))
	d.rd.Lcd.SetCursor(13, 1)
	d.rd.Lcd.Print([]byte(formatMillis(d.rd.Metrics.AcCurrent1Ma, "A")))
	d.rd.Lcd.SetCursor(3, 2)
	d.rd.Lcd.Print([]byte(formatMillis(d.rd.Metrics.AcCurrent2Ma, "A")))
	d.rd.Lcd.SetCursor(13, 2)
	d.rd.Lcd.Print([]byte(formatMillis(d.rd.Metrics.AcCurrent3Ma, "A")))

}

func (d *DisplayManager) screenConf(ev int) {
	d.rd.Lcd.SetCursor(0, 0)
	d.rd.Lcd.Print([]byte("CONF"))
	d.currentScreen = SCREEN_CONF
}

func (d *DisplayManager) screenVersion(ev int) {
	println("screenVersion")
	d.rd.Lcd.ClearDisplay()
	d.rd.Lcd.Print([]byte("VERS"))

	d.rd.Lcd.SetCursor(1, 1)
	d.rd.Lcd.Print([]byte("ZEROCONSO v0.1\n"))
	d.rd.Lcd.Print([]byte("www.oflabs.com\n"))

	d.currentScreen = SCREEN_VERS

}
