package model

import "time"

type Entry struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	AccountId uint      `json:"account_id"`
	Amount    uint      `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func (entry *Entry) TableName() string {
	return "entry"
}
