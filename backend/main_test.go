package main_test

import (
	"backend_ersa/controller"
	"backend_ersa/repository"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Main", func() {
	var (
		mockUserRepo   repository.UserRepository
		userController controller.UserController
		router         *gin.Engine
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		userController = controller.NewUserController(mockUserRepo)
		router = gin.Default()
		router.POST("/register", userController.Register)
		router.POST("/login", userController.Login)
	})

})
