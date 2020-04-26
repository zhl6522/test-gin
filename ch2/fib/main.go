package main

import "fmt"

func main() {
	var a int = 1
	var tmp int
	fmt.Print(a, "")
	for i:=0;i<5;i++ {
		a,tmp = tmp+a,a
		fmt.Print(" ", a)
	}
	fmt.Println()
}
