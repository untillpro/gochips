package errs

import (
	"errors"
	"strings"
)

// Errors is multiple errors container
type Errors []error

func (e Errors) Error() string {
	if len(e) == 1 {
		return e[0].Error()
	}

	var sb strings.Builder
	sb.WriteString("Multiple errors:")
	for _, err := range e {
		sb.WriteString("\n" + err.Error())
	}
	return sb.String()
}

// Add appends new error using provided message
func (e *Errors) Add(message string) Errors {
	*e = append(*e, errors.New(message))
	return *e
}

// AddE appends the provided error
func (e *Errors) AddE(err error) Errors {
	if err == nil {
		return nil
	}
	*e = append(*e, err)
	return *e
}
