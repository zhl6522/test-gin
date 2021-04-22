package main

import (
	"fmt"
	"test-gin/factory/model"
)

// 工厂模式
func main() {
	var stu = model.Student{
		Name:  "zhl",
		Score: 88,
	}
	fmt.Println(stu)
	stu2 := model.Newstudent("mumu", 88.8)
	fmt.Println(*stu2, stu2, stu2.Name, stu2.GetScore())	//不能直接stu2.score调取
}
