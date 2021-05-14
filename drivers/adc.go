package drivers

import (
	"device/stm32"
	"time"
)

const (
	ADC_VREF_MV = 3300
	ADC_MAX     = 4095
)

/*
 *  From: https://github.com/getoffmyhack/STM32F103-Bare-Metal/blob/master/06_ADC/src/main.c
 *  https://www.st.com/resource/en/application_note/cd00211314-how-to-get-the-best-adc-accuracy-in-stm32-microcontrollers-stmicroelectronics.pdf
 *  http://libopencm3.org/docs/latest/stm32f1/html/adc__common__v1_8c_source.html#l00465
 */
func AdcInit() {

	stm32.RCC.CFGR.SetBits(stm32.RCC_CFGR_ADCPRE_Div6) // ADC Clock must be <14Mhz

	stm32.RCC.APB2ENR.SetBits(stm32.RCC_APB2ENR_ADC1EN) // enable ADC clock
	stm32.RCC.APB2ENR.SetBits(stm32.RCC_APB2ENR_IOPAEN) // enable GPIOA clock

	// Switch PA0, PA1, PA2, PA3 to analog mode for ADC Use
	// reset MODE and CNF bits to 0000; MODE = input : CNF = analog mode
	stm32.GPIOA.CRL.ClearBits(0xFF)

	// set sample time for ch 0 to 28.5 cycles (0b011) for PA0 PA1 PA2 PA3
	stm32.ADC1.SMPR2.ReplaceBits(stm32.ADC_SMPR2_SMP0_Cycles28_5, 0b111, 0)
	stm32.ADC1.SMPR2.ReplaceBits(stm32.ADC_SMPR2_SMP0_Cycles28_5, 0b111, 3)
	stm32.ADC1.SMPR2.ReplaceBits(stm32.ADC_SMPR2_SMP0_Cycles28_5, 0b111, 6)
	stm32.ADC1.SMPR2.ReplaceBits(stm32.ADC_SMPR2_SMP0_Cycles28_5, 0b111, 9)

	// put ADC1 into continuous mode and turn on ADC
	stm32.ADC1.CR2.SetBits(stm32.ADC_CR2_CONT | stm32.ADC_CR2_ADON)

	// reset calibration registers
	stm32.ADC1.CR2.SetBits(stm32.ADC_CR2_RSTCAL)

	// wait for calibration register initalized
	for stm32.ADC1.CR2.HasBits(stm32.ADC_CR2_RSTCAL) {
	}

	// enable calibration
	stm32.ADC1.CR2.SetBits(stm32.ADC_CR2_CAL)

	// wait for calibration register initalized
	for stm32.ADC1.CR2.HasBits(stm32.ADC_CR2_CAL) {
	}

	// Conversion starts when ADON bit is set for a second time by software after ADC power-up time
	stm32.ADC1.CR2.SetBits(stm32.ADC_CR2_ADON)

}

// Set current channel to read
// CH0 -> PA0, CH1 -> PA1 ... so on
func AdcSetChannel(channel uint32) {
	stm32.ADC1.SQR3.ReplaceBits(channel, 0b1111, 0)
}

func AdcRead(numSample uint16) uint16 {
	sum := uint64(0)
	for i := uint16(0); i < numSample; i++ {
		sum += uint64(stm32.ADC1.DR.Get() & 0x0FFF)
	}
	return uint16(sum / uint64(numSample))
}

func AcAmplitudeMv() uint32 {
	min := uint32(0xFFFFFFFF)
	max := uint32(0x0)
	timeStart := time.Now()
	cntmes := 0
	for {
		delta := time.Now().Sub(timeStart)
		if delta > time.Millisecond*100 {
			break
		}
		v := uint32(AdcRead(1024))
		cntmes++
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	delta := max - min
	deltaMv := (ADC_VREF_MV * delta) / ADC_MAX
	return deltaMv
}
