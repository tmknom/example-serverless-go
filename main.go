package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler() (string, error) {
	log.Print("Start logging")

	svc := cloudtrail.New(session.New())
	input := &cloudtrail.StartLoggingInput{
		Name: aws.String("default"),
	}
	result, err := svc.StartLogging(input)

	if err != nil {
		log.Printf("Error: %s", err.Error())
		return "failed", err
	}
	return result.GoString(), nil
}

func main() {
	lambda.Start(Handler)
}
