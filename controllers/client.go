package controllers

import(
	"github.com/gin-gonic/gin"
	"../database"
)

type Story struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"body"`
	UserId int `json:"userid"`
	Username string `json:"username"`
}

func IndexData(c *gin.Context){
	db := database.DBConn()
	rows, err := db.Query(
		"SELECT posts.id, posts.title, posts.body, users.id as userid, users.username " + 
		"FROM rivendell.posts, rivendell.users " +
		"WHERE posts.user_id = users.id",
	)
		if err != nil {
			panic(err.Error())
		}
		story := Story{} 
		response := []Story{}

		for rows.Next() {
			var id, userid int
			var title, body, username string

			err = rows.Scan(&id, &title, &body, &userid, &username)
			if err != nil {
				panic(err.Error())
			}

			story.Id = id
			story.Title = title
			story.Content = body
			story.UserId = userid
			story.Username = username
			
			response = append(response, story)
		}
		
		c.JSON(200, response)
	
	defer db.Close()
}


func StoryDetails(c * gin.Context){

	db := database.DBConn()
	rows, err := db.Query("SELECT posts.id, title, body, username, user_id as userid " +
		"FROM rivendell.posts, rivendell.users WHERE posts.id = " + c.Param("id") + 
		" AND posts.user_id = users.id")
	if err != nil{
		c.JSON(500, gin.H{
			"messages" : "Story not found",
		});
	}

	story := Story{}

	for rows.Next(){
		var id, userid int
		var title, body, username string

		err = rows.Scan(&id, &title, &body, &username, &userid)
		if err != nil {
			panic(err.Error())
		}

		story.Id = id
		story.Title = title
		story.Content = body
		story.Username = username
		story.UserId = userid
	}

	c.JSON(200, story)

	defer db.Close()
}

func Categories(c * gin.Context){
	type Category struct{
		Id int `json:"id"`
		Name string `json:"name"`
	}
	db := database.DBConn()
	rows, err := db.Query("SELECT * FROM rivendell.types")
	if err != nil{
		panic(err.Error())
	}
	response := []Category{}
	category := Category{}
	for rows.Next(){
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		category.Id = id
		category.Name = name
		response = append(response, category)
	}
	c.JSON(200, response)
	defer db.Close()
}

func TopFourStories(c * gin.Context){

	type TopFourtItem struct{
		Id int `json:"id"`
		Title string `json:"title"`
		Body string	`json:"body"`
		Views int `json:"views"`
	}

	db := database.DBConn()
	rows, err := db.Query("SELECT id, title, body, views FROM rivendell.posts ORDER BY posts.views desc LIMIT 4")
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
	response := []TopFourtItem{}
	post := TopFourtItem{}
	
	for rows.Next(){
		var id, views int
		var title, body string

		err := rows.Scan(&id, &title, &body, &views)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
		post.Id = id
		post.Title = title
		post.Body = body
		post.Views = views
		response = append(response, post)
	}

	c.JSON(200, response)
	defer db.Close()
}

func TopFiveStories(c * gin.Context) {
	type TopFiveItem struct{
		Id int `json:"id"`
		Title string `json:"title"`
		Body string	`json:"body"`
		Views int `json:"views"`
	}

	db :=  database.DBConn()
	rows, err :=  db.Query("SELECT id, title, body, views FROM rivendell.posts ORDER BY RAND() LIMIT 4")
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	response := []TopFiveItem{}
	post := TopFiveItem{}
	for rows.Next(){
		var id, views int
		var title, body string

		err := rows.Scan(&id, &title, &body, &views)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
		post.Id = id
		post.Title = title
		post.Body = body
		post.Views = views
		response = append(response, post)
	}

	c.JSON(200, response)
	defer db.Close()
}


