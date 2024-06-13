package database

import (
    "lambda-func/types"
)

const DEVICE_TABLE = "closi_devices"

type DeviceStore interface {
    InsertDevice(device types.Device) error
    GetDevice(deviceID string) (types.Device, error)
}
