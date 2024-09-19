package repository

import (
	"backend_ersa/model"

	"gorm.io/gorm"
)

type OrderRepository interface {
	FindByUserID(userID uint) ([]model.Order, error)
	FindByID(id uint) ([]model.Order, error)
	Create(order *model.Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) FindByUserID(userID uint) ([]model.Order, error) {
	var orders []model.Order
	if err := r.db.Preload("OrderItems").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *orderRepository) FindByID(id uint) ([]model.Order, error) {
	var orders []model.Order
	if err := r.db.Preload("OrderItems").Where("id = ?", id).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *orderRepository) Create(order *model.Order) error {
	return r.db.Create(order).Error
}
