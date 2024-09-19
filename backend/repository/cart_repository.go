package repository

import (
	"backend_ersa/model"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindByUserID(userID uint) ([]model.CartItem, error)
	FindByID(id uint) (*model.CartItem, error)
	Create(item *model.CartItem) error
	Update(item *model.CartItem) error
	Delete(id uint) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) FindByUserID(userID uint) ([]model.CartItem, error) {
	var cartItems []model.CartItem
	// result := r.db.Where("user_id = ?", userID).Find(&cartItems)
	// return cartItems, result.Error

	// Fetch cart items along with their associated products
	result := r.db.Preload("Product").Where("user_id = ?", userID).Find(&cartItems)
	return cartItems, result.Error

}

func (r *cartRepository) FindByID(id uint) (*model.CartItem, error) {
	var cartItem model.CartItem
	result := r.db.First(&cartItem, id)
	return &cartItem, result.Error
}

func (r *cartRepository) Create(item *model.CartItem) error {
	return r.db.Create(item).Error
}

func (r *cartRepository) Update(item *model.CartItem) error {
	return r.db.Save(item).Error
}

func (r *cartRepository) Delete(id uint) error {
	return r.db.Delete(&model.CartItem{}, id).Error
}
