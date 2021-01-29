package main

import "fmt"

import "github.com/aws/aws-lambda-go/lambda"

func main() {
	lambdaHandler := NewLambdaHandler()
	lambda.Start(lambdaHandler.handler)

	fmt.Println("{{cookiecutter.directory_name}}, up and running!")
}
