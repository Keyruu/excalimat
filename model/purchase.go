package model

import "gorm.io/gorm"

type Purchase struct {
	gorm.Model
	AccountID uint    `json:"-"`
	Account   Account `json:"account"`
	ProductID uint    `json:"-"`
	Product   Product `json:"product"`
	PaidPrice float32 `json:"paidPrice"`
	Refunded  bool    `json:"refunded"`
}
