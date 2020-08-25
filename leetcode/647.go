package main

import (
	"fmt"
	"time"
)

func main() {
	nums := "aabbbaaaabbbaa"
	st := time.Now()
	len := strings(nums)
	elapsed := time.Since(st)
	fmt.Println("App elapsed: ", elapsed)
	fmt.Println(len)
}

func strings(nums string) int {
	u := 0
	for i := 0; i <= len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			ret := nums[i:j]
			let := len(ret)
			if let/2 == 0 {
				u++
				fmt.Println(ret)
				continue
			}
			for k := 0; k < let/2; k++ {
				nk := let - k - 1
				if ret[k] == ret[nk] {
					fmt.Printf("%v & %v\n", ret[k], ret[nk])
					k2 := k + 1
					if k2 == let/2 {
						u++
						fmt.Println(ret)
					}
				} else {
					break
				}
			}
		}
	}
	return u
}
