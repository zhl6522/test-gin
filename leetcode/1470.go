package main

import "fmt"

func main() {
	nums := []int{2,5,1,3,4,7}
	n := 3
	ret := make([]int, len(nums))
	copy(ret, nums)
	u:=0
	for i := 0; i < len(nums)/2; i++ {
		ret[u] = nums[i]
		u++
		j := n+i
		ret[u] = nums[j]
		u++

	}
	fmt.Println(ret)
}

func main2() {
	nums := []int{2,5,1,3,4,7}
	n := 3
	ret := make([]int, len(nums))
	copy(ret, nums)
	u:=0
	for i := 0; i < len(nums)-1; i+=2 {
		j := i+1
		k := i
		if i > 0 {
			k = i-u
		}
		l := k+n
		ret[i] = nums[k]
		ret[j] = nums[l]
		u++
	}
	fmt.Println(nums,ret)
}
