package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"./controllers"
)



func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"X-Requested-With, X-HTTP-Method-Override, Content-Type, Accept, X-XSRF-TOKEN"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	client := r.Group("/api")
	{
		client.GET("/index/stories/main", controllers.IndexData)
		client.GET("/index/story/:id", controllers.StoryDetails)
		client.GET("/index/categories", controllers.Categories)
		client.GET("/index/stories/top-five", controllers.TopFourStories)
	}
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
