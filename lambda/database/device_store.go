package database

import (
	"fmt"
	"lambda-func/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const DEVICE_TABLE = "closi_devices"

type DeviceStore interface {
	DoesDeviceExist(deviceid string) (bool, error)
	InsertDevice(device types.Device) error
	// GetDevice(deviceID string) (types.Device, error)
}

func (u DynamoDBClient) DoesDeviceExist(deviceid string) (bool, error) {
	result, err := u.databaseStore.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(DEVICE_TABLE),
		Key: map[string]*dynamodb.AttributeValue{
			"deviceid": {
				S: aws.String(deviceid),
			},
		},
	})

	if err != nil {
		return true, err
	}

	if result.Item == nil {
		return false, nil
	}

	return true, nil
}

func (u DynamoDBClient) InsertDevice(device types.Device) error {
	item := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"deviceid": {
				S: aws.String(device.DeviceID),
			},
			"devicetype": {
				S: aws.String(device.DeviceType),
			},
			"datejoined": {
				N: aws.String(fmt.Sprintf("%d", device.DateJoined)),
			},
			"sessionid": {
				S: aws.String(device.SessionID),
			},
		},
	}

	_, err := u.databaseStore.PutItem(item)
	if err != nil {
		return err
	}

	return nil
}
