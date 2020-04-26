package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Print("1111")
}

func TestExchange(t *testing.T)  {
	a := 1
	b := 2
	a,b=b,a
	t.Log(a,b)
}