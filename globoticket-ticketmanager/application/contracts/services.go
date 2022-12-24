package contracts

const (
	LifeTimeSingleton = iota
	LifeTimeSingletonScoped
	LifeTimeTransient
)

type LifeTime int

type ServiceProvider interface {
	GetServiceByName(name string) (any, error)
}

type serviceBuilder func(ServiceProvider) (any, error)

type ServiceCollection interface {
	ServiceProvider
	AddService(name string, service serviceBuilder, lifeTime LifeTime)
	Build() ServiceProvider
}
