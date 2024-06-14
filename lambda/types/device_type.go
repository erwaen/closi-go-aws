package types

import "time"

type RegisterDevice struct {
	DeviceID   string `json:"deviceid"`
	DeviceType string `json:"devicetype"`
}

type Device struct {
	DeviceID   string `json:"deviceid"`
	DeviceType string `json:"devicetype"`
	DateJoined int64  `json:"datejoined"`
	SessionID  string `json:"sessionid"`
}

func NewDevice(registerDevice RegisterDevice) Device {
	return Device{
		DeviceID:   registerDevice.DeviceID,
		DeviceType: registerDevice.DeviceType,
		DateJoined: time.Now().Unix(),
	}
}
