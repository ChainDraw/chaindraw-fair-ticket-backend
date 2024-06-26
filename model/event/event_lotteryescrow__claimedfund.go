package model

import (
	"time"
)

type EventLotteryescrowClaimedfund struct {
	Id         int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	ConcertId  string    `gorm:"column:concert_id;NOT NULL"`
	TicketType string    `gorm:"column:ticket_type;NOT NULL"`
	Organizer  string    `gorm:"column:organizer;NOT NULL"`
	Winner     string    `gorm:"column:winner;NOT NULL"`
	Money      string    `gorm:"column:money;NOT NULL"`
	CreatedAt  time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt  time.Time `gorm:"column:updated_at;NOT NULL"`
}

func (e *EventLotteryescrowClaimedfund) TableName() string {
	return "event_lotteryescrow__claimedfund"
}
