package api

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"lambda-func/database"
	"lambda-func/types"
	"net/http"
)

type DeviceApiHandler struct {
	dbStore database.DeviceStore
}

func NewDeviceApiHandler(dbStore database.DeviceStore) DeviceApiHandler {
	return DeviceApiHandler{
		dbStore: dbStore,
	}
}

func (api DeviceApiHandler) RegisterDeviceHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var registerDevice types.RegisterDevice

	err := json.Unmarshal([]byte(request.Body), &registerDevice)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid Request on unmarshal json",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	if registerDevice.DeviceID == "" || registerDevice.DeviceType == "" {
		return events.APIGatewayProxyResponse{
			Body:       "Please fill device id && device type",
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	deviceExists, err := api.dbStore.DoesDeviceExist(registerDevice.DeviceID)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal server error on finding if Device exist",
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	if deviceExists {
		return events.APIGatewayProxyResponse{
			Body:       "Device already exists",
			StatusCode: http.StatusConflict,
		}, nil
	}

	device := types.NewDevice(registerDevice)

	err = api.dbStore.InsertDevice(device)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal server error in inserting device in store",
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       "Successfully registered device",
		StatusCode: http.StatusOK,
	}, nil
}
