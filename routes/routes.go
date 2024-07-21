package routes

import (
	"inventory-system/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/items", controllers.GetItems)
	r.GET("/items/:id", controllers.GetItem)
	r.POST("/items", controllers.CreateItem)
	r.PUT("/items/:id", controllers.UpdateItem)
	r.DELETE("/items/:id", controllers.DeleteItem)
	return r
}
