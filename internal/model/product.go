package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name    string  `json:"name"`
	Picture string  `json:"picture"`
	Price   float32 `json:"price"`
}
