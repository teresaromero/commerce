package router

import (
	"ecommerce-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine, jwtSecret []byte) {

	// signup, login, and logout
	authRouter := engine.Group("/auth")
	protectedAuthRouter := authRouter.Group("/", middleware.Authenticated(jwtSecret))
	authRouter.POST("/login", nil)
	authRouter.POST("/signup", nil)
	protectedAuthRouter.GET("/logout", nil)

	// Add, remove, and update products in the cart
	cartRouter := engine.Group("/cart")
	protectedCartRouter := cartRouter.Group("/", middleware.Authenticated(jwtSecret))
	protectedCartRouter.GET("/", nil)
	protectedCartRouter.POST("/add", nil)
	protectedCartRouter.DELETE("/remove", nil)
	protectedCartRouter.PUT("/update", nil)

	// Checkout and pay
	checkoutRouter := engine.Group("/checkout")
	protectedCheckoutRouter := checkoutRouter.Group("/", middleware.Authenticated(jwtSecret))
	protectedCheckoutRouter.POST("/", nil)
	protectedCheckoutRouter.GET("/summary", nil)

	// Search and view products
	productRouter := engine.Group("/products")
	productRouter.GET("/", nil)
	productRouter.GET("/:id", nil)
}
