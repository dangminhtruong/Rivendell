package controllers

import(
	"github.com/gin-gonic/gin"
	"../database"
	"../helpers"
	"database/sql"

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

func SignUp(c * gin.Context) {
	type ReqData struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req ReqData
	db := database.DBConn()

	if err := c.ShouldBindJSON(&req); err == nil {
		insUser, insErr := db.Prepare("INSERT INTO users(username, password, token) VALUES(?,?,?)")
		if insErr != nil {
			c.JSON(500, gin.H{
				"messages" : insErr,
			})
		}
		insUser.Exec(req.Username, req.Password, helpers.CreateToken(req.Username, req.Password)) 
		c.JSON(200, gin.H{
			"messages" : "Successfull",
		})
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	
	defer db.Close()
}

func SignIn(c * gin.Context){
	type ReqData struct{
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var token string
	var req ReqData
	db := database.DBConn()

	if err := c.ShouldBindJSON(&req); err == nil {
		rows :=  db.QueryRow("SELECT token FROM users where username='" +
							req.Username + "' and password='" + req.Password + "'")
		reErr := rows.Scan(&token)
		if reErr != nil ||  reErr == sql.ErrNoRows || token == ""{
			c.JSON(200, 500)
		}else{
			c.JSON(200, token)
		}	
	}

	defer db.Close()
}