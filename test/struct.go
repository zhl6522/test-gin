package main

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
)

// 结构体嵌套

type address struct {
	province	string
	city		string
}
type workPlace struct {
	province	string
	city		string
}

type copany struct {
	name	string
	address			//匿名嵌套结构体，等同于：address address
	workPlace
}

// 结构体模拟实现其他语言中的"继承"
type animal struct {
	name	string
}
// 给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%v会动！\n", a.name)
}

// 狗类
type dog struct {
	feet	uint8
	animal			// animal拥有的方法，dog此时也有了
}
// 给dog实现一个会叫的方法
func (d dog) wang() {
	fmt.Printf("%s在叫：汪汪汪~\n", d.name)
}


// 结构体与json

// 1、序列化   json_encode 把Go语言中的结构体变量		-->	json格式的字符串
// 2、反序列化 json_decode json格式的字符串			--> Go语言中能够识别的结构体变量
type person struct {
	Name	string `json:"name" db:"name" ini:"name"`	// 加上`json:"name"`后 {\"Name\":\"李明\",\"Age\":28} 变成小写 {\"name\":\"李明\",\"age\":28}
	Age		int `json:"age"`
	gender	string		//首字母小写只能包内调用，包外输出为空
}

func main() {
	fmt.Println("---------------结构体模拟实现'继承'---------------")
	c1 := copany{
		name: "zhl.icn",
		address: address{
			"安徽",
				"阜阳",
		},
	}
	fmt.Println(c1, c1.address.city)
	//fmt.Println(c1.city)		// 先在自己结构体找这个字段，找不到就去匿名嵌套的结构体中查找该字段		匿名嵌套结构体的字段冲突后，就不能简写了
	fmt.Println(c1.address.city, c1.workPlace.city)

	d1 := dog{
		feet:   4,
		animal: animal{
			"扒扒",
		},
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()

	fmt.Println("---------------结构体与json---------------")
	// 序列化
	p1 := person{
		Name: "李明",
		Age:  28,
		gender:"男",
	}
	b, err := json.Marshal(p1)		//在json的Marshal包中，把p1的内容拿出来转换成字符串，所以gender不会被打印出来
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("%v\n", string(b))
	// 反序列化
	str := `{"name":"小明","age":28}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)	// 传值针是为了能在json.Unmarshal内部修改p2的值
	fmt.Printf("%#v\n", p2)
}
