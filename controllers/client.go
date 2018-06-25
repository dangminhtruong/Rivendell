package controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"../database"
)

func IndexData(c *gin.Context){
	res := ""
	db := database.DBConn()
	rows, err := db.Query("SELECT title, body FROM rivendell.posts")
		if err != nil {
			panic(err.Error())
		}else{
			for rows.Next() {
                var title, body string
				err = rows.Scan(&title, &body)
				if err != nil {
					panic(err.Error())
				}
				res += title + body
			}
			c.String(200, res)
	}
	defer db.Close()
}

func JsonResponse(c *gin.Context){
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "Lena"
	msg.Message = "hey"
	msg.Number = 123
	// Note that msg.Name becomes "user" in the JSON
	// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
	c.JSON(http.StatusOK, msg)
}

/* func PostDetails(c *gin.Context){
	type post struct {
		Title string
		Body string
	}

	type response struct {
		Item post
		Relative []post
	}

	

	c.JSON(http.StatusOK, res)
} */