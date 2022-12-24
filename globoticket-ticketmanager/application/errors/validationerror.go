package errors

import (
	"strings"
)

type validationError struct {
	message          string
	validationErrors map[string]string
}

type ValidationError interface {
	AddValidationError(property string, errorMessage string) ValidationError
}

func (e *validationError) NewValidationError(message string) ValidationError {
	return &validationError{message: message}
}

func (e *validationError) AddValidationError(property string, errorMessage string) ValidationError {
	e.validationErrors[property] = errorMessage
	return e
}

// use errors.As(err, &ValidationError{}) -- using var ve ValidationError

func (e *validationError) Error() string {
	b := strings.Builder{}
	b.WriteString(e.message)
	b.WriteString("\n")
	for _, e := range e.validationErrors {
		b.WriteString("\t")
		b.WriteString(e)
		b.WriteString("\n")
	}
	return b.String()
}
