package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args)>1 {
		fmt.Println("hello word", os.Args[1])
	}
}
