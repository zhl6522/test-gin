package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	ip := "192.168.10.125"
	st := time.Now()
	str := defangIPaddr(ip)
	elapsed := time.Since(st)
	fmt.Println("App elapsed: ", elapsed)
	fmt.Println(str)
}

func defangIPaddr(ip string) string {
	return strings.ReplaceAll(ip, ".", "[.]")
}
