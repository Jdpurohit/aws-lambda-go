package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type App struct {
	id string
}

func newApp(id string) *App {
	return &App{
		id: id,
	}
}

func (app *App) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	responseBody := map[string]string{
		"message": "Hi, All looks good!",
	}

	responseJSON, err := json.Marshal(responseBody)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       `{"error": "internal server error"}`,
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type":                     "text/plain",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Headers":     "Content-Type",
			"Access-Control-Allow-Methods":     "POST, OPTIONS, GET, PUT",
			"Access-Control-Allow-Credentials": "true",
		},
		Body: string(responseJSON),
	}, nil
}

func main() {
	id := "golang-app"
	app := newApp(id)
	lambda.Start(app.Handler)
}
