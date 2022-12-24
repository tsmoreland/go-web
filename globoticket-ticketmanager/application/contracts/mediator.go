package contracts

import "reflect"

type MediatorRequest interface {
}

type MediatorRequestHandler interface {
	CanHandle(req MediatorRequest)
	Handle(req MediatorRequest) chan MediatorResponse
}

type MediatorResponse interface {
	Error() error
	HasResponse() bool
	Response() (chan any, error)
}

type Mediator interface {
	Send(req MediatorRequest) MediatorResponse
}

type MediatorOptions interface {
	AddHandler(requestType reflect.Type, handler MediatorRequestHandler)
}
