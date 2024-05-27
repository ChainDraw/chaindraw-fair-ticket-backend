package model

import (
	"time"
)

type EventNftListed struct {
	Id             int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Seller         string    `gorm:"column:seller;NOT NULL"`
	LotteryAddress string    `gorm:"column:lottery_address;NOT NULL"`
	TokenId        string    `gorm:"column:token_id;NOT NULL"`
	Price          string    `gorm:"column:price;NOT NULL"`
	CreatedAt      time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt      time.Time `gorm:"column:updated_at;NOT NULL"`
}

func (e *EventNftListed) TableName() string {
	return "event_nft_listed"
}
