package controllers

import(
	"github.com/gin-gonic/gin"
	"../database"
)

type Story struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"body"`
}

func IndexData(c *gin.Context){
	db := database.DBConn()
	rows, err := db.Query("SELECT id, title, body FROM rivendell.posts")
		if err != nil {
			panic(err.Error())
		}
		story := Story{} 
		response := []Story{}

		for rows.Next() {
			var id int
			var title, body string

			err = rows.Scan(&id, &title, &body)
			if err != nil {
				panic(err.Error())
			}

			story.Id = id
			story.Title = title
			story.Content = body
			
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