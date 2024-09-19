package controller

import (
	"backend_ersa/model"
	"backend_ersa/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	GetAllCategories(c *gin.Context)
	GetCategoryByID(c *gin.Context)
	AddCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
}

type categoryController struct {
	repo repository.CategoryRepository
}

func NewCategoryController(categoryRepository repository.CategoryRepository) CategoryController {
	return &categoryController{repo: categoryRepository}
}

func (s *categoryController) GetAllCategories(c *gin.Context) {
	categories, err := s.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch all product categories"})
		return

	}

	c.JSON(http.StatusOK, categories)
}

func (s *categoryController) GetCategoryByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Category ID"})
		return
	}

	product, err := s.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (s *categoryController) AddCategory(c *gin.Context) {
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.repo.Create(&category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add category"})
		return
	}

	c.JSON(http.StatusCreated, category)
}

func (s *categoryController) UpdateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category item ID"})
		return
	}

	var updateData model.Category
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingItem, err := s.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category item not found"})
		return
	}

	existingItem.Name = updateData.Name
	existingItem.Description = updateData.Description

	err = s.repo.Update(existingItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category item"})
		return
	}

	c.JSON(http.StatusOK, existingItem)
}

func (s *categoryController) DeleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category item ID"})
		return
	}

	fmt.Printf("Logging from controller: %d", id)

	err = s.repo.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "category removed successfully"})
}

// func GetAllCategory(c *gin.Context) {
// 	var categories []model.Category
// 	result := database.DB.Find(&categories)
// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching categories"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, categories)
// }

// func GetCategoryById(c *gin.Context) {
// 	id := c.Param("id")
// 	var categories model.Category
// 	result := database.DB.First(&categories, id)
// 	if result.Error != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, categories)
// }
