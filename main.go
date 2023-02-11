package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kubil-ismail/go-example/controllers"
	"github.com/kubil-ismail/go-example/models"
)

func main() {
	route := gin.Default()

	models.ConnectDb()

	route.GET("/api/product", controllers.Index)
	route.GET("/api/product/:id", controllers.Detail)
	route.POST("/api/product", controllers.Add)
	route.PATCH("/api/product/:id", controllers.Update)
	route.DELETE("/api/product/:id", controllers.Delete)

	route.Run()
}
