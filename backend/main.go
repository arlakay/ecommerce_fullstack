package main

import (
	"backend_ersa/controller"
	"backend_ersa/database"
	"backend_ersa/middleware"
	"backend_ersa/model"
	"backend_ersa/repository"
	"backend_ersa/route"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	db := database.NewDB()
	dbCredential := database.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: "ecommerce",
		Port:         5432,
	}

	dbConn, err := db.Connect(&dbCredential)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	dbConn.AutoMigrate(&model.User{}, &model.Product{}, &model.Category{}, &model.ProductCategory{}, &model.CartItem{}, &model.Order{}, &model.OrderItem{})

	userRepo := repository.NewUserRepo(dbConn)
	userController := controller.NewUserController(userRepo)

	productRepo := repository.NewProductRepo(dbConn)
	productController := controller.NewProductController(productRepo)

	categoryRepo := repository.NewCategoryRepo(dbConn)
	categoryController := controller.NewCategoryController(categoryRepo)

	cartRepo := repository.NewCartRepository(dbConn)
	cartController := controller.NewCartController(cartRepo)

	orderRepo := repository.NewOrderRepository(dbConn)
	orderController := controller.NewOrderController(orderRepo)

	route.SetupRoutes(r, userController, productController, categoryController, cartController, orderController)

	r.Run(":8080")
}
