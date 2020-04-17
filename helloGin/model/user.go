package model

type User struct {
	Name    string `json:"name" form:"name"`
	Address string `json:"address" form:"address"`
	Age     string `json:"age" form:"age"`
	Mobile  string `json:"mobile" form:"mobile"`
	Sex     string `json:"sex" form:"sex"`
	Id      uint16 `json:"id" form:"id"`
}
