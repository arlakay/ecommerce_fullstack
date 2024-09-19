package repository

import (
	"backend_ersa/model"
	"fmt"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll() ([]model.Category, error)
	FindByID(id uint) (*model.Category, error)
	Create(category *model.Category) error
	Update(category *model.Category) error
	Delete(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (u *categoryRepository) FindAll() ([]model.Category, error) {
	var categories []model.Category
	if err := u.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) FindByID(id uint) (*model.Category, error) {
	var category model.Category
	result := r.db.First(&category, id)
	return &category, result.Error
}

func (r *categoryRepository) Create(category *model.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) Update(category *model.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id uint) error {
	fmt.Printf("Logging from repository: %d", id)

	return r.db.Delete(&model.Category{}, id).Error
}
