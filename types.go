package rigctld

import "errors"

// Frequency is a frequency in hertz.
type Frequency int64

// Mode is a radio modulation mode.
type Mode int

const (
	// USB is Upper Sideband.
	USB Mode = iota
	// LSB is Lower Sideband.
	LSB
	// CW is Continuous Wave.
	CW
	// CWR is Continuous Wave Reverse.
	CWR
	// RTTY is Radio Teletype.
	RTTY
	// RTTYR is Radio Teletype Reverse.
	RTTYR
	// AM is Amplitude Modulation.
	AM
	// FM is Frequency Modulation.
	FM
	// WFM is Wideband FM.
	WFM
	// AMS is Amplitude Modulation Synchronous.
	AMS
	// PKTLSB is Packet Lower Sideband.
	PKTLSB
	// PKTUSB is Packet Upper Sideband.
	PKTUSB
	// PKTFM is Packet FM.
	PKTFM
	// ECSSUSB is Exalted Carrier Single Sideband Upper Sideband.
	ECSSUSB
	// ECSSLSB is Exalted Carrier Single Sideband Lower Sideband.
	ECSSLSB
	// FA is Frequency Agile.
	FA
	// SAM is Synchronous AM.
	SAM
	// SAL is Synchronous AM Lower.
	SAL
	// SAH is Synchronous AM Upper.
	SAH
	// DSB is Double Sideband.
	DSB
)

func (m Mode) String() string {
	return [...]string{
		"USB",
		"LSB",
		"CW",
		"CWR",
		"RTTY",
		"RTTYR",
		"AM",
		"FM",
		"WFM",
		"AMS",
		"PKTLSB",
		"PKTUSB",
		"PKTFM",
		"ECSSUSB",
		"ECSSLSB",
		"FA",
		"SAM",
		"SAL",
		"SAH",
		"DSB"}[m]
}

func (m Mode) EnumIndex() int {
	return int(m)
}

// ModeFromString converts a string to a Mode.
func ModeFromString(modeStr string) (Mode, error) {
	modes := map[string]Mode{
		"USB":     USB,
		"LSB":     LSB,
		"CW":      CW,
		"CWR":     CWR,
		"RTTY":    RTTY,
		"RTTYR":   RTTYR,
		"AM":      AM,
		"FM":      FM,
		"WFM":     WFM,
		"AMS":     AMS,
		"PKTLSB":  PKTLSB,
		"PKTUSB":  PKTUSB,
		"PKTFM":   PKTFM,
		"ECSSUSB": ECSSUSB,
		"ECSSLSB": ECSSLSB,
		"FA":      FA,
		"SAM":     SAM,
		"SAL":     SAL,
		"SAH":     SAH,
		"DSB":     DSB,
	}

	mode, exists := modes[modeStr]
	if !exists {
		return 0, errors.New("invalid mode string")
	}
	return mode, nil
}
