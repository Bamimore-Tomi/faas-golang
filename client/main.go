package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/joho/godotenv"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-lambda-go/lambda"
)

type Info struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

type Response struct {
	Profile string `json:"profile"`
}

func main() {
	// lambda.Start(HandleLambdaEvent)
	// event := MyEvent{Name: "Tomi", Age: 17}
	// c, err := json.Marshal(event)
	// if err != nil {
	// 	fmt.Println("Error")
	// }
	godotenv.Load()
	sess := session.Must(session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable}))
	client := lambda.New(sess, &aws.Config{Region: aws.String("us-east-1"), Credentials: credentials.NewStaticCredentials(os.Getenv("aws_access_key_id"), os.Getenv("aws_secret_access_key"), "")})
	request := Info{Firstname: "Oluwatomisin", Lastname: "Bamimore", Age: 16}
	payload, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshalling  request")
		os.Exit(0)
	}
	result, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String("user-profile"), Payload: payload})
	if err != nil {
		fmt.Println("Error calling user-profile function", err)
		os.Exit(0)
	}
	var resp Response
	err = json.Unmarshal(result.Payload, &resp)
	if err != nil {
		fmt.Println("Error unmarshalling user-profile response")
		os.Exit(0)
	}
	fmt.Println(resp.Profile)
}
