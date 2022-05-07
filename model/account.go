package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model `valid:"-"`
	ExtID      string  `json:"extID" valid:"required"`
	Email      string  `json:"email" valid:"email,required"`
	Name       string  `json:"name" valid:"required"`
	Balance    float32 `json:"balance" valid:"numeric,optional"`
	Picture    string  `json:"picture" valid:"optional"`
	PIN        string  `json:"-" valid:"-"`
}
