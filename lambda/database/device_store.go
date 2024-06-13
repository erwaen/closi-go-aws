package database

import (
    "lambda-func/types"
)

const DEVICE_TABLE = "deviceTable"

type DeviceStore interface {
    InsertDevice(device types.Device) error
    GetDevice(deviceID string) (types.Device, error)
}
