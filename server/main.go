package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type InfoEvent struct {
	firstname string
	lastname  string
	age       int
}

type Response struct {
	Profile string `json:"profile"`
}

func HandleInfoEvent(event InfoEvent) (Response, error) {
	return Response{Profile: fmt.Sprintf("Their name is %s %s , they are %d ", event.firstname, event.lastname, event.age)}, nil
}

func main() {
	lambda.Start(HandleInfoEvent)
}
