package shared

import (
	"bytes"
	"fmt"
)

func (info *AllDevices) String() string {
	b := bytes.NewBuffer(nil)

	b.WriteString(fmt.Sprintln("Playback Devices:"))
	for i, d := range info.Playback {
		var defaultStatus string
		if d.IsDefault {
			defaultStatus = "x"
		} else {
			defaultStatus = " "
		}
		b.WriteString(fmt.Sprintf("  %d: [%s] %s\n\t  Status: [%s], Channels: %d-%d, Sample Rate: %d-%d\n",
			i, defaultStatus, d.Name, d.Status, d.MinChannels, d.MaxChannels, d.MinSampleRate, d.MaxSampleRate))
	}

	b.WriteString(fmt.Sprintln(""))
	b.WriteString(fmt.Sprintln("Capture Devices:"))
	for i, d := range info.Capture {
		var defaultStatus string
		if d.IsDefault {
			defaultStatus = "x"
		} else {
			defaultStatus = " "
		}
		b.WriteString(fmt.Sprintf("  %d: [%s] %s\n\t  Status: [%s], Channels: %d-%d, Sample Rate: %d-%d\n",
			i, defaultStatus, d.Name, d.Status, d.MinChannels, d.MaxChannels, d.MinSampleRate, d.MaxSampleRate))
	}
	return b.String()
}
