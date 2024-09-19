package controller

import (
	"backend_ersa/model"
	"backend_ersa/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	GetUserOrders(c *gin.Context)
	CreateOrder(c *gin.Context)
	GetOrderDetails(c *gin.Context)
}

type orderController struct {
	repo repository.OrderRepository
}

func NewOrderController(orderRepository repository.OrderRepository) OrderController {
	return &orderController{repo: orderRepository}
}

func (s *orderController) GetUserOrders(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	orders, err := s.repo.FindByUserID(uint(userID))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch user orders"})
		return
	}

	c.JSON(200, orders)
}

func (s *orderController) CreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := s.repo.Create(&order)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(201, order)
}

func (s *orderController) GetOrderDetails(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := s.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(200, order)
}
