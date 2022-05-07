package model

import "gorm.io/gorm"

type Purchase struct {
	gorm.Model `valid:"-"`
	AccountID  uint    `json:"-" valid:"required"`
	Account    Account `json:"account" valid:"-"`
	ProductID  uint    `json:"-" valid:"required"`
	Product    Product `json:"product" valid:"-"`
	PaidPrice  float32 `json:"paidPrice" valid:"required"`
	Refunded   bool    `json:"refunded" valid:"required"`
}
