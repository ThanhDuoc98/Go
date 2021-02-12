package model

import "time"

type Transfer struct {
	Id            uint      `json:"id" gorm:"primaryKey"`
	FromAccountId uint      `json:"from_account_id"`
	ToAccountId   uint      `json:"to_account_id"`
	Amount        uint      `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}

func (transfer *Transfer) TableName() string {
	return "transfer"
}
