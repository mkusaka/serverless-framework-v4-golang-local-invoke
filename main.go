package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string      `json:"message"`
	Input   interface{} `json:"input"`
}

func handler(ctx context.Context, request Request) (Response, error) {
	if request.Name == "fail" {
		return Response{}, errors.New("failed")
	}
	fmt.Printf("Received request: %+v\n", request)
	fmt.Fprintf(os.Stderr, "Received request: %+v\n", request)

	response := Response{
		Message: fmt.Sprintf("Hello, %s!", request.Name),
		Input:   request,
	}

	responseJSON, _ := json.MarshalIndent(response, "", "  ")
	fmt.Printf("Sending response: %s\n", string(responseJSON))
	fmt.Fprintf(os.Stderr, "Sending response: %s\n", string(responseJSON))

	return response, nil
}

func main() {
	startMsg := "Lambda function starting..."
	fmt.Println(startMsg)
	fmt.Fprintf(os.Stderr, "%s\n", startMsg)
	lambda.Start(handler)
}
