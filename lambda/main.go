package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"lambda-func/app"
	"lambda-func/middleware"
	"net/http"
)

type MyEvent struct {
	Username string `json:"username"`
}

// Take in a payload and do something with it
func HandleRequest(event MyEvent) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}

	return fmt.Sprintf("Successfully called by - %s", event.Username), nil
}

func ProtectedHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "This is a secret path",
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	myApp := app.NewApp()
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch request.Path {
		case "/register":
			return myApp.ApiHandler.User.RegisterUserHandler(request)
		case "/login":
			return myApp.ApiHandler.User.LoginUser(request)
		case "/protected":
			return middleware.ValidateJWTMiddleware(ProtectedHandler)(request)
		case "/registerdevice":
			return myApp.ApiHandler.Device.RegisterDeviceHandler(request)
		default:
			return events.APIGatewayProxyResponse{
				Body:       "Not found",
				StatusCode: http.StatusNotFound,
			}, nil
		}
	})

}
