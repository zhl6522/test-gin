package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"helloGin/database"
	"helloGin/model"
	"log"
	"net/http"
)

//var db *sql.DB

func init() {
	db = database.GetDataBase()
}

func QueryParam(c *gin.Context) {
	var u model.User
	row, err := db.Query("select id,`name`,address,age,mobile,sex from user")
	if err != nil {
		fmt.Println(err)
	}
	defer row.Close()

	users := make([]model.User, 0)

	for row.Next() {
		err := row.Scan(&u.Id, &u.Name, &u.Address, &u.Age, &u.Mobile, &u.Sex)
		users = append(users, u)
		if err != nil {
			fmt.Println(err)
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"result": users,
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	//name := c.Query("name")	//url?name=zhl
	name := c.Param("name")
	var u model.User
	row := db.QueryRow("select id,`name`,address,age,mobile,sex from user where id = ? and `name` = ?", id, name)
	err := row.Scan(&u.Id, &u.Name, &u.Address, &u.Age, &u.Mobile, &u.Sex)
	checkError(err)
	c.JSON(http.StatusOK, gin.H{
		"result": u,
	})
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	res, err := db.Exec("delete from user  where id=?", id)
	var count int64
	count, err = res.RowsAffected()
	checkError(err)
	if count != 1 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}

func RenderView(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.HTML(http.StatusOK, "./template/hello.html", gin.H{
		"title": "first",
	})
}
