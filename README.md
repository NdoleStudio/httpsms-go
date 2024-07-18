# httpsms-go

[![Build](https://github.com/NdoleStudio/httpsms-go/actions/workflows/main.yml/badge.svg)](https://github.com/NdoleStudio/httpsms-go/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/NdoleStudio/httpsms-go/branch/main/graph/badge.svg)](https://codecov.io/gh/NdoleStudio/httpsms-go)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/NdoleStudio/httpsms-go/badges/quality-score.png?b=main)](https://scrutinizer-ci.com/g/NdoleStudio/httpsms-go/?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/NdoleStudio/httpsms-go)](https://goreportcard.com/report/github.com/NdoleStudio/httpsms-go)
[![GitHub contributors](https://img.shields.io/github/contributors/NdoleStudio/httpsms-go)](https://github.com/NdoleStudio/httpsms-go/graphs/contributors)
[![GitHub license](https://img.shields.io/github/license/NdoleStudio/httpsms-go?color=brightgreen)](https://github.com/NdoleStudio/httpsms-go/blob/master/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/NdoleStudio/httpsms-go)](https://pkg.go.dev/github.com/NdoleStudio/httpsms-go)


This package provides a generic `go` client template for the Http SMS Api

## Installation

`httpsms-go` is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/NdoleStudio/httpsms-go
```

Alternatively the same can be achieved if you use `import` in a package:

```go
import "github.com/NdoleStudio/httpsms-go"
```


## Implemented

- [x] **[Messages](#messages)**
  - [x] `POST /v1/messages/send`: Send a new SMS
  - [x] `GET /v1/messages`: Get list of messages which are exchanged between 2 phone numbers.
- [x] **Heartbeats**
  - [x] `GET /v1/heartbeats`: Get the heartbeats of an Android Phone
- [x] **Message Threads**
  - [x] `GET /v1/message-threads`: Get the message threads of a phone number
  - [x] `DELETE v1/message-threads/:messageThreadID`: Delete a message thread
- [x] **Cipher**
  - [x] `Encrypt`: Encrypt the content of a message to cipher text
  - [x] `Decrypt`: Decrypt an encrypted message content to plain text


## Usage

### Initializing the Client

An instance of the client can be created using `httpsms.New()`.

```go
package main

import (
    "github.com/NdoleStudio/httpsms-go"
)

func main()  {
    client := htpsms.New(htpsms.WithAPIKey(/* API Key from https://httpsms.com/settings */))
}
```

### Error handling

All API calls return an `error` as the last return object. All successful calls will return a `nil` error.

```go
_, response, err := client.Messages.Send(context.Background())
if err != nil {
    //handle error
}
```

### Messages

#### `POST /v1/messages/send`: Send a new SMS Message

```go
message, response, err := client.Messages.Send(context.Background(), &MessageSendParams{
    Content: "This is a sample text message",
    From:    "+18005550199",
    To:      "+18005550100",
})

if err != nil {
    log.Fatal(err)
}

log.Println(message.Code) // 202
```

## Testing

You can run the unit tests for this client from the root directory using the command below:

```bash
go test -v
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
