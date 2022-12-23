package domain

import (
	"github.com/google/uuid"
	"time"
)

type Entities interface {
	AuditableEntity | Event | Category | Order
}

type Event struct {
	AuditableEntity
	EventId     uuid.UUID
	Name        string
	Price       int
	Artist      string
	Date        time.Time
	Description string
	ImageUrl    string
	CategoryId  uuid.UUID
	Category    *Category
}

type Category struct {
	AuditableEntity
	CategoryId uuid.UUID
	Name       string
	Events     []Event
}

type Order struct {
	AuditableEntity
	EventId     uuid.UUID
	UserId      uuid.UUID
	OrderTotal  int
	OrderPlaced time.Time
	OrderPaid   bool
}
