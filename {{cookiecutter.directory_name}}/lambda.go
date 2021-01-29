package main

import "github.com/aws/aws-lambda-go/events"

// LambdaHandler encapsulates the interaction with AWS Lambda.
type LambdaHandler struct {
}

// NewLambdaHandler create a new lambda handler.
func NewLambdaHandler() *LambdaHandler {
	return &LambdaHandler{}
}

func (lh *LambdaHandler) handler(event events.{{cookiecutter.recieve_event}}) (events.{{cookiecutter.return_event}}, error) {
	logger.Info("hi from handler implementation.")
	return event, nil
}
