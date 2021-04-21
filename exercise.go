package main

import "fmt"

func main() {
	var score [3][5]int
	for i := 0; i < len(score); i++ {
		for j := 0; j < len(score[i]); j++ {
			fmt.Printf("请输入第%d班第%d位同学的成绩：", i+1, j+1)
			fmt.Scanln(&score[i][j])
		}
	}
	var all = 0
	var length = 0
	for i := 0; i < len(score); i++ {
		var s = 0
		for j := 0; j < len(score[i]); j++ {
			s += score[i][j]
		}
		all += s
		length+= len(score[i])
		avg := s / len(score[i])
		fmt.Printf("第%d班的平均成绩：%d\n", i+1, avg)
	}
	allavg := all / length
	fmt.Printf("所有班级的平均成绩：%d\n", allavg)
	//fmt.Println(score)
}
