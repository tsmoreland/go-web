package geteventslist

import "github.com/tsmoreland/go-web/globoticket-ticketmanager/application/contracts"

type Query struct {
}

func NewGetEventsListRequest() contracts.MediatorRequest {
	return &Query{}
}
