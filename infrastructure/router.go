package infrastructure

import (
	"APIHEX_LPySP/infrastructure/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handlers.UserHandler, productHandler *handlers.ProductHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/users", userHandler.GetUsers)
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	r.GET("/products", productHandler.GetProducts)
	r.POST("/products", productHandler.CreateProduct)
	r.GET("/products/:id", productHandler.GetProductByID)
	r.PUT("/products/:id", productHandler.UpdateProduct)
	r.DELETE("/products/:id", productHandler.DeleteProduct)

	return r
}
