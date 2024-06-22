package domain

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusReserved  SpotStatus = "reserved"
)

type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketID string
}