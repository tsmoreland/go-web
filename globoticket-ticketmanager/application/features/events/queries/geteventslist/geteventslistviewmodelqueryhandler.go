package geteventslist

import (
	"github.com/tsmoreland/go-web/globoticket-ticketmanager/application/contracts"
	"github.com/tsmoreland/go-web/globoticket-ticketmanager/application/errors"
	"reflect"
)

var (
	queryType = reflect.TypeOf(&Query{})
)

type QueryHandler struct {
	repository contracts.EventsRepository
}

func (handler *QueryHandler) CanHandle(req contracts.MediatorRequest) bool {
	return reflect.TypeOf(req) == reflect.TypeOf(queryType)
}

func (handler *QueryHandler) Handle(req contracts.MediatorRequest) (chan contracts.MediatorResponse, error) {
	if !handler.CanHandle(req) {
		return nil, errors.NewBadRequestError("unsupported request type")
	}
	r := make(chan contracts.MediatorResponse)
	go func() {
		channel := handler.repository.GetAll()
		items := <-channel
		if items.Err != nil {
			r <- NewResponseFromError(items.Err)
		} else {
			r <- NewResponse(items.Items)
		}
	}()
	return r, nil

}
