package controllers

import(
	"github.com/gin-gonic/gin"
	"../database"
)

func CreateNewPost(c * gin.Context){
	db := database.DBConn()
	
	type CreatePost struct {
		Title     string `form:"title" json:"title" binding:"required"`
		Body string `form:"body" json:"body" binding:"required"`
		Status int `form:"status" json:"status" binding:"required"`
		TypeId int `form:"typeId" json:"typeId" binding:"required"`
		UserId int `form:"userId" json:"userId" binding:"required"`
	}

	var json CreatePost


	if err := c.ShouldBindJSON(&json); err == nil {
		insPost, err := db.Prepare("INSERT INTO rivendell.posts(title, body, status, type_id, avata, user_id) VALUES(?,?,?,?,?,?)",)
		if err != nil {
			c.JSON(200, gin.H{
				"messages" : err,
			})
		}
		insPost.Exec(json.Title, json.Body, json.Status, json.TypeId, "ava1.jpg", json.UserId) 
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	defer db.Close()
}