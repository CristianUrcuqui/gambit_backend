package main

import (
	"context"
	"gambit_backend/awsgo"
	"gambit_backend/bd"
	"gambit_backend/handlers"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(cxt context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	awsgo.StartedAws()

	if !ValidateParameter() {
		panic("Error getting parameter: send parameter 'SecretName', 'UrlPrefix' ")
	}

	var res *events.APIGatewayProxyResponse
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	bd.ReadSecret()

	status, message := handlers.Handlers(path, method, body, header, request)

	headerResp := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headerResp,
	}

	return res, nil

}

func ValidateParameter() bool {
	_, bringParams := os.LookupEnv("SecretName")
	if !bringParams {
		return bringParams
	}
	_, bringParams = os.LookupEnv("UrlPrefix")
	if !bringParams {
		return bringParams
	}

	return bringParams
}
