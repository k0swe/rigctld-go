package rigctld

import (
	"errors"
	"fmt"
)

func (s *Client) GetFreq() (Frequency, error) {
	resp, err := s.writeRead("f\n")
	if err != nil {
		return 0, err
	}

	var freq Frequency
	_, err = fmt.Sscanf(resp, "%d\n", &freq)
	if err != nil {
		return 0, err
	}
	return freq, nil
}

func (s *Client) SetFreq(freq Frequency) error {
	resp, err := s.writeRead(fmt.Sprintf("F %d\n", freq))
	if err != nil {
		return err
	}

	var report int
	_, err = fmt.Sscanf(resp, "RPRT %d\n", &report)
	if err != nil {
		return err
	}
	if report != 0 {
		// TODO: Can this be more specific? Probably would have to delve into the rigctld source
		return errors.New(fmt.Sprintf("rigctld error %d", report))
	}
	return nil
}

// GetMode returns the current mode and bandpass width of the radio.
func (s *Client) GetMode() (Mode, Frequency, error) {
	resp, err := s.writeRead("m\n")
	if err != nil {
		return -1, 0, err
	}

	var modeStr string
	var bandpass Frequency
	_, err = fmt.Sscanf(resp, "%s\n%d\n", &modeStr, &bandpass)
	if err != nil {
		return -1, 0, err
	}
	mode, err := ModeFromString(modeStr)
	if err != nil {
		return -1, 0, err
	}
	return mode, bandpass, nil
}

func (s *Client) SetMode(mode Mode, bandpass Frequency) error {
	modeStr := mode.String()
	resp, err := s.writeRead(fmt.Sprintf("M %s %d\n", modeStr, bandpass))
	if err != nil {
		return err
	}

	var report int
	_, err = fmt.Sscanf(resp, "RPRT %d\n", &report)
	if err != nil {
		return err
	}
	if report != 0 {
		// TODO: Can this be more specific? Probably would have to delve into the rigctld source
		return errors.New(fmt.Sprintf("rigctld error %d", report))
	}
	return nil
}
