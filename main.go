package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"./controllers"
)



func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"X-Requested-With, X-HTTP-Method-Override, Content-Type, Accept, X-XSRF-TOKEN"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	client := r.Group("/client")
	{
		client.GET("/stories/main", controllers.IndexData)
		client.GET("/story/:id", controllers.StoryDetails)
		client.GET("/categories", controllers.Categories)
		client.GET("/stories/top-four", controllers.TopFourStories)
		client.GET("/stories/random", controllers.TopFiveStories)
	}

	admin := r.Group("/admin")
	{
		admin.POST("/story/create", controllers.CreateNewPost)
	}

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
