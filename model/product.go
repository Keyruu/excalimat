package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string  `json:"name" validate:"required"`
	Picture    string  `json:"picture" validate:""`
	Price      float32 `json:"price" validate:"required" gorm:"type:decimal(9,2);"`
	BundleSize int8    `json:"bundleSize" validate:"required"`
	Type       string  `json:"type" validate:"required"`
}
