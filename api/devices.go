package api

import (
	"github.com/JustinTimperio/streaMaster/shared"
	"github.com/gen2brain/malgo"
)

func GetDevices() (ad shared.AllDevices, err error) {
	var (
		in  []shared.DeviceInfo
		out []shared.DeviceInfo
	)

	// Connect to miniaudio via malgo
	context, err := malgo.InitContext(nil, malgo.ContextConfig{}, nil)
	if err != nil {
		return shared.AllDevices{Playback: out, Capture: in}, err
	}

	defer func() {
		_ = context.Uninit()
		context.Free()
	}()

	// Fetch info on playback devices
	playbackDevices, err := context.Devices(malgo.Playback)
	if err != nil {
		return shared.AllDevices{Playback: out, Capture: in}, err
	}

	for _, info := range playbackDevices {
		var (
			pDevice shared.DeviceInfo
			e       = "ok"
		)

		full, err := context.DeviceInfo(malgo.Playback, info.ID, malgo.Shared)
		if err != nil {
			e = err.Error()
		}

		if full.IsDefault == 1 {
			pDevice.IsDefault = true
		} else {
			pDevice.IsDefault = false
		}

		pDevice.Status = e
		pDevice.Name = info.Name()
		pDevice.MinChannels = full.MinChannels
		pDevice.MaxChannels = full.MaxChannels
		pDevice.MinSampleRate = full.MinSampleRate
		pDevice.MaxSampleRate = full.MaxSampleRate

		out = append(out, pDevice)
	}

	// Fetch info on capture devices
	captureDevices, err := context.Devices(malgo.Capture)
	if err != nil {
		return shared.AllDevices{Playback: out, Capture: in}, err
	}

	for _, info := range captureDevices {
		var (
			cDevice shared.DeviceInfo
			e       = "ok"
		)

		full, err := context.DeviceInfo(malgo.Capture, info.ID, malgo.Shared)
		if err != nil {
			e = err.Error()
		}

		if full.IsDefault == 1 {
			cDevice.IsDefault = true
		} else {
			cDevice.IsDefault = false
		}

		cDevice.Status = e
		cDevice.Name = info.Name()
		cDevice.MinChannels = full.MinChannels
		cDevice.MaxChannels = full.MaxChannels
		cDevice.MinSampleRate = full.MinSampleRate
		cDevice.MaxSampleRate = full.MaxSampleRate

		in = append(in, cDevice)

	}
	return shared.AllDevices{Playback: out, Capture: in}, nil
}
