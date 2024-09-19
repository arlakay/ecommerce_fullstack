package controller

import (
	"backend_ersa/model"
	"backend_ersa/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartController interface {
	GetCart(c *gin.Context)
	AddToCart(c *gin.Context)
	UpdateCartItem(c *gin.Context)
	RemoveFromCart(c *gin.Context)
}

type cartController struct {
	repo repository.CartRepository
}

func NewCartController(cartRepository repository.CartRepository) CartController {
	return &cartController{repo: cartRepository}
}

func (s *cartController) GetCart(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	cartItems, err := s.repo.FindByUserID(uint(userID))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch cart items"})
		return
	}

	c.JSON(200, cartItems)
}

func (s *cartController) AddToCart(c *gin.Context) {
	var cartItem model.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := s.repo.Create(&cartItem)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to add item to cart"})
		return
	}

	c.JSON(201, cartItem)
}

func (s *cartController) UpdateCartItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid cart item ID"})
		return
	}

	var updateData model.CartItem
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	existingItem, err := s.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Cart item not found"})
		return
	}

	existingItem.Quantity = updateData.Quantity

	err = s.repo.Update(existingItem)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update cart item"})
		return
	}

	c.JSON(200, existingItem)
}

func (s *cartController) RemoveFromCart(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid cart item ID"})
		return
	}

	err = s.repo.Delete(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to remove item from cart"})
		return
	}

	c.JSON(200, gin.H{"message": "Item removed from cart successfully"})
}
