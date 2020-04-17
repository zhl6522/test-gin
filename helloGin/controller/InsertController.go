package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"helloGin/database"
	"helloGin/model"
	"net/http"
)

var db *sql.DB

func init() {
	db = database.GetDataBase()
}

func InsertUser(c *gin.Context) {
	var u model.User
	err := c.ShouldBindJSON(&u)
	res, err := db.Exec("insert into user (`name`,address,age,mobile,sex) value (?,?,?,?,?)",
		&u.Name, &u.Address, &u.Age, &u.Mobile, &u.Sex)
	var count int64
	count, err = res.RowsAffected()
	checkError(err)
	if count != 1 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
		})
	} else {
		/*c.JSON(http.StatusOK, gin.H{
			"success":true,
		})*/
		c.Redirect(http.StatusMovedPermanently, "/hello")
	}
}

func UpdateUser(c *gin.Context) {
	var u model.User
	err := c.ShouldBindJSON(&u)
	res, err := db.Exec("update user set `name`=?,address=?,age=?,mobile=?,sex=? where id=?",
		&u.Name, &u.Address, &u.Age, &u.Mobile, &u.Sex, &u.Id)
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
