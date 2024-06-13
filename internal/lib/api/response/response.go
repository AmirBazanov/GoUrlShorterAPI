package response

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "ERROR"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("%s is required", err.Field()))
		case "url":
			errMsgs = append(errMsgs, fmt.Sprintf("Invalid URL: %s", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("Invalid field: %s", err.Tag()))
		}
	}
	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, "; "),
	}
}
