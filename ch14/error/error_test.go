package error__test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2")
var LargetThanTwoError = errors.New("n should be not larget than 100")
func GetFibonacci(n int) ([]int, error) {
	//if n < 2 || n > 100 {
	//	return nil,errors.New("n should be in [2,100]")
	//}
	if n < 2 {
		return nil,LessThanTwoError
	}
	if n > 100 {
		return nil,LargetThanTwoError
	}
	fibList := []int{1,1}
	for i:=2;i<n;i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v,err := GetFibonacci(15); err != nil {
		if err == LessThanTwoError {
			fmt.Println("It is less.")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func GetFibonacci1(str string) {
	var (
		i int
		err error
		list []int
	)
	if i,err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Error", err)
	}
}

//避免嵌套
func GetFibonacci2(str string) {
	var (
		i int
		err error
		list []int
	)
	if i,err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}
	if list,err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)
}