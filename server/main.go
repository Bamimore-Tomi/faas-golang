package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type InfoEvent struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

type Response struct {
	Profile string `json:"profile"`
}

func HandleInfoEvent(event InfoEvent) (Response, error) {
	return Response{Profile: fmt.Sprintf("Their name is %s %s, they are %d ", event.Firstname, event.Lastname, event.Age)}, nil
}

func main() {
	lambda.Start(HandleInfoEvent)
}
