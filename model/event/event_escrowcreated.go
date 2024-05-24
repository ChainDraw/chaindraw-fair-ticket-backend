package model

import (
	"time"
)

type EventEscrowCreated struct {
	Id            int32     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	ConcertId     string    `gorm:"column:concert_id;NOT NULL"`
	TicketType    string    `gorm:"column:ticket_type;NOT NULL"`
	EscrowAddress string    `gorm:"column:escrow_address;NOT NULL"`
	CreatedAt     time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt     time.Time `gorm:"column:updated_at;NOT NULL"`
}

func (e *EventEscrowCreated) TableName() string {
	return "event_escrow_created"
}
