package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// function to handle GET method
// @return response object to proxy back through API Gateway
func HandleGet() events.APIGatewayProxyResponse {
	// init API Gateway response struct with headers, status, body
	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: http.StatusOK,
		Body:       "service operational",
	}
}

// function to handle POST method
// @param ctx is context from Lambda invocation
// @param req is request object handling incoming API webhook
// @return response object to proxy back through API Gateway
// @return error
func HandlePost(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) (
	events.APIGatewayProxyResponse,
	error,
) {
	// init API Gateway response struct with headers
	res := events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}

	// init db
	// dynamoSvc := db.New()

	return res, nil
}

// function to handle Lambda event
// @param ctx is context from Lambda invocation
// @param req is request object handling incoming API webhook
// @return response object to proxy back through API Gateway
// @return error
func HandleLambdaEvent(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) (
	events.APIGatewayProxyResponse,
	error,
) {
	// define variable to hold API Gateway response object
	var res events.APIGatewayProxyResponse

	// logic to manage varying request method
	switch {

	case req.HTTPMethod == "POST":
		// handle POST request
		res, err := HandlePost(
			ctx,
			req,
		)
		if err != nil {
			return res, err
		}
	case req.HTTPMethod == "GET":
		// handle GET request
		res = HandleGet()
	}

	return res, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
