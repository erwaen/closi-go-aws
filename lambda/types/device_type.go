package types

import "time"

type RegisterDevice struct {
	DeviceID string `json:"deviceid"`
}

type Device struct {
	DeviceID   string `json:"deviceid"`
	DateJoined int64  `json:"datejoined"`
	SessionID  string `json:"sessionid"`
}

func NewDevice(registerDevice RegisterDevice) Device {
	return Device{
		DeviceID:   registerDevice.DeviceID,
		DateJoined: time.Now().Unix(),
	}
}
