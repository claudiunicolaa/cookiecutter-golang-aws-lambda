package main

import "fmt"

import "github.com/aws/aws-lambda-go/lambda"
import "go.uber.org/zap"

var (
	logger, _ = zap.NewProduction()
)

func main() {
	lambdaHandler := NewLambdaHandler()
	lambda.Start(lambdaHandler.handler)

	fmt.Println("{{cookiecutter.directory_name}}, up and running!")
}
