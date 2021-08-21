package main

import (
	"fmt"
)

type InfoEvent struct {
	Firstname string
	Lastname  string
	Age       int
}

type Response struct {
	Profile string `json:"profile"`
}

func HandleInfoEvent(event InfoEvent) (Response, error) {
	return Response{Profile: fmt.Sprintf("Their name is %s %s , they are %d ", event.Firstname, event.Lastname, event.Age)}, nil
}

func main() {
	request := InfoEvent{Firstname: "Oluwatomisin", Lastname: "Bamimore", Age: 16}
	fmt.Println(HandleInfoEvent(request))
	// lambda.Start(HandleInfoEvent)
}
