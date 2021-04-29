package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

//声明一个StuSlice结构体切片类型
type StuSlice []Student

func (stu StuSlice) Len() int {
	return len(stu)
}

func (stu StuSlice) Less(i, j int) bool {
	return stu[i].Score > stu[j].Score
}

func (stu StuSlice) Swap(i, j int) {
	stu[i], stu[j] = stu[j], stu[i]
}

//接口是对继承的补充
//接口编译的经典案例	对复杂切片排序
func main() {
	var stues StuSlice
	for i := 0; i < 10; i++ {
		str := Student{
			Name:  fmt.Sprintf("学生：%d", rand.Intn(100)),
			Age:   rand.Intn(100),
			Score: rand.Intn(150),
		}
		stues = append(stues, str)
	}
	fmt.Println("排序前----------")
	for _,v := range stues{
		fmt.Println(v)
	}
	sort.Sort(stues)
	fmt.Println("排序后----------")
	for _,v := range stues{
		fmt.Println(v)
	}
}
