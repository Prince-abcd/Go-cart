package database

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
type Item struct {
	gorm.Model
	Name  string
	Price uint
}
type User struct {
	gorm.Model
	Username string
	Password string
}

type Cart struct {
	gorm.Model
	UserID uint
	Items  []Item `gorm:"many2many:cart_items;"`
}

type Order struct {
	gorm.Model
	UserID uint
	Items  []Item `gorm:"many2many:order_items;"`
}
