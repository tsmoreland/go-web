package contracts

import (
	"github.com/google/uuid"
	"github.com/tsmoreland/go-web/globoticket-ticketmanager/domain"
)

type Repository[T domain.Entities] interface {
	GetById(uuid uuid.UUID) chan SingleResponse[T]
	GetAll() chan MultipleResponse[T]
	GetPage(pageNumber int, pageSize int) chan PageResponse[T]
	Add(entity T) chan SingleResponse[T]
	Update(entity T) chan error
	Delete(entity T) chan error
}

type EventsRepository interface {
	Repository[domain.Event]
}

type CategoriesRepository interface {
	Repository[domain.Category]
}

type OrdersRepository interface {
	Repository[domain.Order]
}

type SingleResponse[T domain.Entities] struct {
	Item T
	Err  error
}

type MultipleResponse[T domain.Entities] struct {
	Items []T
	Err   error
}

type PageResponse[T domain.Entities] struct {
	PageNumber int
	PageSize   int
	TotalPages int
	TotalCount int
	Items      []T
	Err        error
}
