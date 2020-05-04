package array_test_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1,2,3,4}
	arr3 := [...]int{1,3,4,5}
	arr1[1] = 5
	t.Log(arr[0],arr[1])
	t.Log(arr1,arr3)
}

func TestArrayTravel(t *testing.T) {
	arr2 := [...]int{1,3,5,7}
	/*for i:=0;i<len(arr2);i++ {
		t.Log(arr2[i])
	}*/
	for _,e := range arr2 {
		t.Log(e)
	}
	for idx,e := range arr2 {
		t.Log(idx, e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [2][2]int{{1,3},{2,4}}
	t.Log(arr3,arr3[1:])
	arr4 := [...]int{1,2,3,4,5}
	t.Log(arr4[1:2],arr4[1:3],arr4[2:len(arr4)],arr4[2:],arr4[:3])		//不支持 arr4[:-1]、arr4[-1:]
}
