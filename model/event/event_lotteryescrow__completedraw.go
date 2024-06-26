package model

import (
	"time"
)

type EventLotteryescrowCompletedraw struct {
	Id             int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	LotteryAddress string    `gorm:"column:lottery_address;NOT NULL"`
	CreatedAt      time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt      time.Time `gorm:"column:updated_at;NOT NULL"`
}

func (e *EventLotteryescrowCompletedraw) TableName() string {
	return "event_lotteryescrow__completedraw"
}
