package contracts

import "reflect"

type MediatorRequest interface {
}

type MediatorRequestHandler interface {
	CanHandle(req MediatorRequest)
	Handle(req MediatorRequest) chan MediatorResponse
}

type MediatorRequestHandlerFactory interface {
	Build() MediatorRequestHandler
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
	AddHandler(requestType reflect.Type, handlerFactory MediatorRequestHandlerFactory)
}
