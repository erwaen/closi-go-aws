package api

import (
	"lambda-func/database"
)

type ApiHandler struct {
	User   UserApiHandler
	Device DeviceApiHandler
}

func NewApiHandler(db database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		User:   NewUserApiHandler(db),
		Device: NewDeviceApiHandler(db),
	}
}
