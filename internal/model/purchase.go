package model

import "gorm.io/gorm"

type Purchase struct {
	gorm.Model
	Account   Account `gorm:"foreignKey:ID"`
	Product   Product `gorm:"foreignKey:ID"`
	PaidPrice float32 `json:"paidPrice"`
}
