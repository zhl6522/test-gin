package main

import (
	"fmt"
	"test-gin/test/split_string"
)

func main() {
	ret := split_string.Split("babcbdef", "b")
	fmt.Printf("%#v\n", ret)
}
