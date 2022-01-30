package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string  `json:"name"`
	Picture    string  `json:"picture"`
	Price      float32 `json:"price"`
	BundleSize int8    `json:"bundleSize"`
	Type       string  `json:"type"`
	Archived   bool    `json:"archived"`
}
