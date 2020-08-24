package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	n := []int{4, 5, null, null, 2, 1}
	st := time.Now()
	str := minDepth(n)
	elapsed := time.Since(st)
	fmt.Println("App elapsed: ", elapsed)
	fmt.Println(str)
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}
	return minD + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
