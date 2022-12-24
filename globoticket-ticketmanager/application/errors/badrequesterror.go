package errors

type badRequestError struct {
	message string
}

func NewBadRequestError(message string) error {
	return &badRequestError{message: message}
}

func (e *badRequestError) Error() string {
	return e.message
}
