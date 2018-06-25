package main

import (
	"github.com/gin-gonic/gin"
	"./controllers"
)



func setupRouter() *gin.Engine {
	r := gin.Default()

	client := r.Group("/api/client")
	{
		client.GET("/index-data", controllers.IndexData)
		//client.GET("/post/:id", controllers.PostDetails)
	}
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
