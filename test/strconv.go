package main

import "fmt"

func main() {
	i := int(97)
	ret1 := string(i)	//"97"
	fmt.Println(ret1)
	ret2 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v",ret2)

}
