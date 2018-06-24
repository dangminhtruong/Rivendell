package main

import (
	"github.com/gin-gonic/gin"
	"./controllers"
)



func setupRouter() *gin.Engine {
	r := gin.Default()

	client := r.Group("/api/client")
	{
		client.GET("/index-data", controllers.Getme)
		client.GET("/contact-data", controllers.JsonResponse)
	}
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
