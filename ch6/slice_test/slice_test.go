package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1,2,3,4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
			/*[]type, len, cap
			其中len个元素会被初始化为默认零值，未初始化元素不可访问
			*/
	t.Log(len(s2), cap(s2))
	t.Log(s2[0],s2[1],s2[2])
	s2 = append(s2, 1,2)
	t.Log(s2[0],s2[1],s2[2],s2[3],s2[4])
	t.Log(len(s2), cap(s2))
	s2 = append(s2, 3)
	t.Log(s2[0],s2[1],s2[2],s2[3],s2[4],s2[5])
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s2 := []int{}
	for i:=0;i<10;i++ {
		s2 = append(s2, i)
		t.Log(len(s2), cap(s2), s2)
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sept","Oct","Nov","Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(Q2, year)
}
/*
					数组 VS 切片

1、容量是否可伸缩	否		是
2、是否可以进行比较	是		否
*/