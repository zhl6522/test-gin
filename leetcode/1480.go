package main

import (
	"fmt"
)

func main() {
	nums := []int{1,2,3,4,5,6,7}
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i]+nums[i-1]
	}
	fmt.Println(nums)
}
