package main

import "fmt"

func main() {
	k := 0
	nums := []int{1,2,2,3,2,1}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if(nums[i] == nums[j]) {
				fmt.Println(i,j)
				k ++
			}
		}
	}
	fmt.Println(k)
}
