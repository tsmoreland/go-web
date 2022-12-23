package errors

import "fmt"

type notFoundError struct {
	id         any
	entityType any
}

func (e *notFoundError) NewNotFoundError(id any, entityType any) error {
	return &notFoundError{id: id, entityType: entityType}
}

func (e *notFoundError) Error() string {
	return fmt.Sprintf("%v matching %v not found", e.entityType, e.id)
}
