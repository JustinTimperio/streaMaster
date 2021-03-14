package shared

type AllDevices struct {
	Playback []DeviceInfo
	Capture  []DeviceInfo
}

type DeviceInfo struct {
	Name          string
	Status        string
	MinChannels   uint32
	MaxChannels   uint32
	MinSampleRate uint32
	MaxSampleRate uint32
	IsDefault     bool
}
