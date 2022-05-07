package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model `valid:"-"`
	Name       string  `json:"name" valid:"required"`
	Picture    string  `json:"picture" valid:"required"`
	Price      float32 `json:"price" valid:"required"`
	BundleSize int8    `json:"bundleSize" valid:"required"`
	Type       string  `json:"type" valid:"required"`
	Archived   bool    `json:"archived" valid:"required"`
}
