package route

import (
	"backend_ersa/controller"
	"backend_ersa/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, uc controller.UserController, pc controller.ProductController, cc controller.CategoryController, cartC controller.CartController, oc controller.OrderController) {
	// Public routes
	r.POST("/register", uc.Register)
	r.POST("/login", uc.Login)

	r.GET("/products", pc.GetAllProducts)
	r.GET("/products/:id", pc.GetProductByID)
	r.GET("/products/category/:category_id", pc.GetProductByCategory)

	r.GET("/categories", cc.GetAllCategories)
	r.GET("/categories/:id", cc.GetCategoryByID)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{

		protected.GET("/users", uc.UsersAll)
		protected.GET("/user", uc.UserByToken)
		protected.GET("/users/:id", uc.UserByID)

		protected.POST("/categories", cc.AddCategory)
		protected.PUT("/categories/:id", cc.UpdateCategory)
		protected.DELETE("/categories/:id", cc.DeleteCategory)

		protected.POST("/products", pc.AddProduct)
		protected.PUT("/products/:id", pc.UpdateProduct)
		protected.DELETE("/products/:id", pc.DeleteProduct)

		protected.GET("/cart/user/:user_id", cartC.GetCart)
		protected.POST("/cart", cartC.AddToCart)
		protected.PUT("/cart/:id", cartC.UpdateCartItem)
		protected.DELETE("/cart/:id", cartC.RemoveFromCart)

		protected.GET("/orders/user/:user_id", oc.GetUserOrders)
		protected.POST("/orders", oc.CreateOrder)
		protected.GET("/orders/:id", oc.GetOrderDetails)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
	})
}
