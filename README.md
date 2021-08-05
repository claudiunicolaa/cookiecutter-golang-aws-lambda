# cookiecutter-golang-aws-lambda

## Generate project
```bash
cookiecutter https://github.com/claudiunicolaa/cookiecutter-golang-aws-lambda
```

## Structure

```bash

{{cookiecutter.directory_name}}
├── README.md
├── go.mod
├── go.sum
├── lambda.go
└── main.go

```

## Template items
```json
{
    "directory_name": "the directory name for your lambda",
    "module_path": "module path will be added on go.mod",
    "recieve_event": "the type of event will be recieved",
    "return_event": "the type of event will be return"
}
```

| :memo:        |   possible values of `recieve_event` and `return_event` are listed [here](https://pkg.go.dev/github.com/aws/aws-lambda-go/events)    |
|---------------|:------------------------|
