package main

import (
	"fmt"
	"test-gin/test/split/string"
)

func main() {
	ret := string.Split("babcdef", "b")
	fmt.Printf("%#v\n", ret)
	ret2 := string.Split("bbb", "b")
	fmt.Printf("%#v\n", ret2)
	ret3 := string.Split("acb", "b")
	fmt.Printf("%#v\n", ret3)
}
