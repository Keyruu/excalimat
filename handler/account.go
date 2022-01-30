package handler

import (
	"errors"

	"github.com/keyruu/excalimat-backend/database"
	"github.com/keyruu/excalimat-backend/model"
	"gorm.io/gorm"
)

func getAccountByID(id string) (*model.Account, error) {
	db := database.DB
	var user model.Account
	if err := db.Where(&model.Account{}, id).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
