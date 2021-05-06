package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		d
		e, f = iota, iota 	//4 4 同一行当一次计算
		g    = iota			//5
	)
	fmt.Println(a, b, c, d, e, f, g)
}
