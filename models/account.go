package models

import (
	"gorm.io/gorm"
	//"gorm.io/drivers/postgres"
)

type Account struct {
	gorm.Model
	Name    string `json:"name"`
	Type    string `json:"type"`
	Balance int64  `json:"balance"`
}

func MigrateAccount(db *gorm.DB) error {

	error := db.AutoMigrate(&Account{})
	return error

}
