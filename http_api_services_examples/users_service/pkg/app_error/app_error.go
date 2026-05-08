package app_error

import "fmt"

type Kind int

const (
	KindNotFound Kind = iota
	KindInvalidInput
	KindInternal
	KindAlreadyExists
)

type AppError struct {
	Kind    Kind
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func NewNotFound(msg string) *AppError {
	return &AppError{Kind: KindNotFound, Message: msg}
}

func NewInvalidInput(msg string) *AppError {
	return &AppError{Kind: KindInvalidInput, Message: msg}
}

func NewInternal(err error) *AppError {
	return &AppError{Kind: KindInternal, Message: "internal error", Err: err}
}

func NewAlreadyExists(msg string) *AppError {
	return &AppError{Kind: KindAlreadyExists, Message: msg}
}
