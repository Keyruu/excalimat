package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model `valid:"-"`
	ExtID      string  `json:"extID" validate:"required"`
	Email      string  `json:"email" validate:"email,required"`
	Name       string  `json:"name" validate:"required"`
	Balance    float32 `json:"balance" validate:"" gorm:"type:decimal(9,2);"`
	Picture    string  `json:"picture" validate:""`
	PIN        string  `json:"-" validate:"-"`
}
