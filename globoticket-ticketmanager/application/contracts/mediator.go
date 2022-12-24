package contracts

type MediatorRequest interface {
}

type MediatorRequestHandler interface {
	CanHandle(req MediatorRequest) bool
	Handle(req MediatorRequest) (chan MediatorResponse, error)
	Close() error
}

type MediatorRequestHandlerFactory interface {
	Build(provider ServiceProvider) MediatorRequestHandler
}

type MediatorResponse interface {
	Error() error
	HasResponse() bool
	Response() (any, error)
	Close() error
}

type Mediator interface {
	Send(req MediatorRequest) MediatorResponse
}

type MediatorOptions interface {
	AddHandler(name string, handlerFactory MediatorRequestHandlerFactory)
}
