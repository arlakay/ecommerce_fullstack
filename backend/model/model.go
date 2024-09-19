package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

type User struct {
	BaseModel
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type Category struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Product struct {
	BaseModel
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"image_url"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	// Categories  []Category `json:"categories" gorm:"many2many:product_categories;"`
}

type ProductCategory struct {
	ProductID  int `json:"product_id"`
	CategoryID int `json:"category_id"`
}

type CartItem struct {
	BaseModel
	UserID    uint      `json:"user_id"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	AddedAt   time.Time `json:"added_at" gorm:"default:CURRENT_TIMESTAMP"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID"`
	// User      User `json:"user" gorm:"foreignKey:UserID"`
}

type Order struct {
	BaseModel
	UserId     uint        `json:"user_id"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"`
	Total      float64     `json:"total"`
	Status     string      `json:"status"`
}

type OrderItem struct {
	BaseModel
	OrderId   uint    `json:"order_id"`
	ProductId uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type OrderResult struct {
	OrderID   uint    `json:"order_id"`
	UserID    uint    `json:"user_id"`
	Total     float64 `json:"total"`
	Status    string  `json:"status"`
	ItemID    uint    `json:"order_item_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
