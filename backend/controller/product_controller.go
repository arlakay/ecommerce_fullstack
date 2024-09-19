package controller

import (
	"net/http"
	"strconv"

	"backend_ersa/model"
	"backend_ersa/repository"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	GetAllProducts(c *gin.Context)
	GetProductByID(c *gin.Context)
	GetProductByCategory(c *gin.Context)
	AddProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

type productController struct {
	repo repository.ProductRepository
}

func NewProductController(productRepository repository.ProductRepository) ProductController {
	return &productController{repo: productRepository}
}

func (s *productController) GetAllProducts(c *gin.Context) {
	products, err := s.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch all products"})
		return

	}

	c.JSON(http.StatusOK, products)
}

func (s *productController) GetProductByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
		return
	}

	product, err := s.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (s *productController) GetProductByCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("category_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Category ID"})
		return
	}

	products, err := s.repo.FindByCategoryID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products by category"})
		return
	}

	c.JSON(http.StatusOK, products)

}

func (s *productController) AddProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.repo.Create(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (s *productController) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product item ID"})
		return
	}

	var updateData model.Product
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingItem, err := s.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product item not found"})
		return
	}

	existingItem.Name = updateData.Name
	existingItem.Description = updateData.Description
	existingItem.ImageUrl = updateData.ImageUrl
	existingItem.Price = updateData.Price
	existingItem.Stock = updateData.Stock

	err = s.repo.Update(existingItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product item"})
		return
	}

	c.JSON(http.StatusOK, existingItem)
}

func (s *productController) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product item ID"})
		return
	}

	err = s.repo.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove product item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product item removed successfully"})
}
