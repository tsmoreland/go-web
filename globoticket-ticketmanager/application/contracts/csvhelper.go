package contracts

import (
	"github.com/tsmoreland/go-web/globoticket-ticketmanager/application/features/events/queries/geteventsexport"
)

type CsvExporter interface {
	ExportEventsToCsvAsync([]geteventsexport.EventsExportDto) chan ExportResponse
	ExportEventsToCsv([]geteventsexport.EventsExportDto) ExportResponse
}

type ExportResponse struct {
	Data []byte
	Err  error
}
