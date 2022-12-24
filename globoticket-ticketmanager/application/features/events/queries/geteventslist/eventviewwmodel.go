package geteventslist

import (
	"github.com/google/uuid"
	"time"
)

type EventViewModel struct {
	EventId  uuid.UUID `json:"event_id"`
	Name     string    `json:"name"`
	Date     time.Time `json:"date"`
	ImageUrl *string   `json:"image_url"`
}
