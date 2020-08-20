package main

import (
	"fmt"
)

func main() {
	candies := []int{2,13,5,1,13}
	//extraCandies := 3
	// 获取最大值
	max := candies[0]
	for i := 1; i < len(candies); i++ {
		if candies[i]>max {
			max = candies[i]
		}
	}
	ret := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if candies[i] < max {
			ret[i] = false
		} else {
			ret[i] = true
		}
	}
	fmt.Println(candies,ret)
}
