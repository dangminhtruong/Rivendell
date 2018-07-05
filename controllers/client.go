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
		"where posts.user_id = users.id",
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
	rows, err := db.Query("SELECT id, title, body FROM rivendell.posts where id = " + c.Param("id"))
	if err != nil{
		c.JSON(500, gin.H{
			"messages" : "Story not found",
		});
	}

	story := Story{}

	for rows.Next(){
		var id int
		var title, body string

		err = rows.Scan(&id, &title, &body)
		if err != nil {
			panic(err.Error())
		}

		story.Id = id
		story.Title = title
		story.Content = body
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
	for rows.Next(){
		var id int
		var name string
		category := Category{}
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		category.Id = id
		category.Name = name
		response = append(response, category);
	}
	c.JSON(200, response)
	defer db.Close()
}