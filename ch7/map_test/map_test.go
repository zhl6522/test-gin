package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1:1, 2:4, 3:9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d m2[4]=%d", len(m2), m2[4])
	m3 := make(map[int]int, 10)		//make第二个参数是cap，但是不能用cap(m3)来打印make的cap
	//t.Logf("len m3=%d", cap(m3))	//.\map_test.go:13:25: invalid argument m3 (type map[int]int) for cap
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	//区分默认为0和赋值为0
	if v,ok:=m1[2];ok {
		t.Logf("key 2`s value is %d", v)
	} else {
		t.Log("key 2 is not existing")
	}
}
//map元素的访问
//在访问的key不存在时，仍会返回零值，不能通过返回nil来判断元素是否存在

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1:1, 2:4, 3:9}
	for key,v := range m1{
		t.Log(key,v)
	}
}