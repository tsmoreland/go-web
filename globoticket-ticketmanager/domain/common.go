package domain

import (
	"errors"
	"math"
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

type Page[T Entities] struct {
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
	TotalPages int `json:"total_pages"`
	TotalCount int `json:"total_count"`
	Items      []T `json:"items"`
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

func (req PageRequest) CalculateTotalPages(totalCount int) int {
	if req.PageSize > 0 {
		return int(math.Ceil(float64(totalCount) / float64(req.PageSize)))
	} else {
		return 0
	}
}
