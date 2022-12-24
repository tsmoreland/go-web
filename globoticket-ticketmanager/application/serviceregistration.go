package application

import "github.com/tsmoreland/go-web/globoticket-ticketmanager/application/contracts"

func AddApplicationServices(services contracts.ServiceCollection) contracts.ServiceCollection {
	//mediatorOptions.AddHandler(reflect.TypeOf(geteventslist.Query{}).Name(), nil)
	return services
}
