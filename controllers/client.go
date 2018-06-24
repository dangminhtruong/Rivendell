package controllers

import(
	"github.com/gin-gonic/gin"
	"database/sql"
	"fmt"
	"net/http"
)

func Getme(c *gin.Context){
	res := ""
	db, err := sql.Open("mysql", "root:789852@/civ3")
	if err != nil {
		fmt.Println(err)
		c.String(200, "Connect failure")	
	}else{
		rows, err := db.Query("SELECT username, email FROM civ3.users")
		if err != nil {
			panic(err.Error())
		}else{
			for rows.Next() {
                var username, email string
				err = rows.Scan(&username, &email)
				if err != nil {
					panic(err.Error())
				}
				res += username + email
			}
			c.String(200, res)
		}
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