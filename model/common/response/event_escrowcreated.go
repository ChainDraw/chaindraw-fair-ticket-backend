package commonresp

type EventEscrowCreated struct {
	ConcertId     string `gorm:"column:concert_id;NOT NULL" json:"concert_id"`
	TicketType    string `gorm:"column:ticket_type;NOT NULL" json:"ticket_type"`
	EscrowAddress string `gorm:"column:escrow_address;NOT NULL" json:"escrow_address"`
}

func (e *EventEscrowCreated) TableName() string {
	return "event_escrow_created"
}
