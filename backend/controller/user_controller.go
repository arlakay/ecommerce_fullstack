package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"backend_ersa/model"
	"backend_ersa/repository"
	"backend_ersa/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	UsersAll(c *gin.Context)
	UserByToken(c *gin.Context)
	UserByID(c *gin.Context)

	// CheckPassLength(pass string) bool
	// CheckPassAlphabet(pass string) bool
}

type userController struct {
	repo repository.UserRepository
}

func NewUserController(userRepository repository.UserRepository) UserController {
	return &userController{repo: userRepository}
}

func (s *userController) Login(c *gin.Context) {
	var loginUser model.User
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := s.repo.Login(loginUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := util.GenerateToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	fmt.Printf("Token UserID: %v\n", userID)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (s *userController) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := s.repo.FindByEmail(user.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with this email already exists"})
		return
	}

	err := s.repo.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user_id": user.ID})
}

func (s *userController) UsersAll(c *gin.Context) {
	users, err := s.repo.UsersAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch all users"})
		return

	}

	c.JSON(http.StatusOK, users)
}

func (s *userController) UserByToken(c *gin.Context) {

	// Extract token from the Authorization header
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
		return
	}

	// Remove "Bearer " prefix if it exists
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Parse the JWT token
	claims, err := util.ValidateToken(tokenString)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		}
		c.Abort()
		return
	}

	// Extract user ID from claims
	userID := claims.UserID
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	// Fetch the user using the userID
	user, err := s.repo.UserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Return the user data
	c.JSON(http.StatusOK, user)
}

func (s *userController) UserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	users, err := s.repo.UserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, users)
}
