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
		AllowAllOrigins: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers, X-Requested-With, X-HTTP-Method-Override, Content-Type, Accept, X-XSRF-TOKEN"},
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
		admin.POST("/signup", controllers.SignUp)
		admin.POST("/login", controllers.SignIn)
	}

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
