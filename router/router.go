package router

import (
	"github.com/gin-gonic/gin"
	"product-management-system/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/products", handlers.CreateProduct)
	r.GET("/products/:id", handlers.GetProductByID)
	r.GET("/products", handlers.GetProducts)

	return r
}
