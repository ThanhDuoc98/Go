package model

import "time"

type Account struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Owner     string    `json:"owner"`
	Balance   uint      `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

func (account *Account) TableName() string {
	return "account"
}
