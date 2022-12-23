package domain

import (
	"errors"
	"time"
)

var (
	ErrInvalidPageNumber = errors.New("invalid page number, must be greater than 0")
	ErrInvalidPageSize   = errors.New("invalid page size, must be greater than 0")
)

type AuditableEntity struct {
	CreatedBy        *string
	CreatedDate      *time.Time
	LastModifiedBy   *string
	LastModifiedDate *time.Time
}

type PageRequest struct {
	PageNumber int
	PageSize   int
}

func (req PageRequest) Validate() []error {
	var errs []error
	if req.PageSize <= 0 {
		errs = append(errs, ErrInvalidPageNumber)
	}
	if req.PageNumber <= 0 {
		errs = append(errs, ErrInvalidPageSize)
	}
	return errs
}
