package geteventslist

import (
	"github.com/tsmoreland/go-web/globoticket-ticketmanager/application/contracts"
	"github.com/tsmoreland/go-web/globoticket-ticketmanager/domain"
)

type response struct {
	items []domain.Event
	err   error
}

func NewResponseFromError(err error) contracts.MediatorResponse {
	return &response{items: nil, err: err}
}
func NewResponse(items []domain.Event) contracts.MediatorResponse {
	return &response{items: items, err: nil}
}

func (r *response) Error() error {
	return r.err
}

func (r *response) HasResponse() bool {
	return true
}

func (r *response) Response() (any, error) {
	if r.err != nil {
		return nil, r.err
	} else {
		return r.items, nil
	}
}

func (r *response) Close() error {
	return nil
}
