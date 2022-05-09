package model

import "gorm.io/gorm"

type Purchase struct {
	gorm.Model
	AccountID uint    `json:"accountId" validate:"required"`
	Account   Account `json:"account" validate:"dive"`
	ProductID uint    `json:"productId" validate:"required"`
	Product   Product `json:"product" validate:"dive"`
	PaidPrice float32 `json:"paidPrice" validate:"required" gorm:"type:decimal(9,2);"`
}
