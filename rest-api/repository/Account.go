package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"rest-api/model"
)

func GetAllAccounts(db *gorm.DB, account *[]model.Account) (err error) {
	if err = db.Find(account).Error; err != nil {
		return err
	}
	return nil
}

func GetAccount(db *gorm.DB, account *model.Account, id uint) (err error) {
	if err = db.Where("id = ?", id).First(account).Error; err != nil {
		return err
	}
	return nil
}

func CreateAccount(db *gorm.DB, account *model.Account) (err error) {
	if err = db.Create(account).Error; err != nil {
		return err
	}
	return nil
}
