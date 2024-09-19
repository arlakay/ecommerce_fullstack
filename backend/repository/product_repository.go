package repository

import (
	"backend_ersa/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]model.Product, error)
	FindByID(id uint) (*model.Product, error)
	FindByCategoryID(categoryId uint) ([]model.Product, error)
	Create(product *model.Product) error
	Update(product *model.Product) error
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (u *productRepository) FindAll() ([]model.Product, error) {
	var products []model.Product
	if err := u.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) FindByID(id uint) (*model.Product, error) {
	var product model.Product
	result := r.db.First(&product, id)
	return &product, result.Error
}

func (r *productRepository) FindByCategoryID(categoryID uint) ([]model.Product, error) {
	var products []model.Product
	err := r.db.Joins("JOIN product_categories ON products.id = product_categories.product_id").
		Where("product_categories.category_id = ?", categoryID).
		Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) Update(product *model.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}
