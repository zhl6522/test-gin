package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id		string
	Name	string
	Age		int
}

func TestCreateEmployeeObj(t *testing.T) {
	e:= Employee{"0", "Bobo", 20}
	e1:=Employee{Name:"Benbo", Age:20}
	e2:=new(Employee)
	e2.Id = "2"
	e2.Name = "Uu"
	e2.Age = 18
	t.Log(e)
	t.Log(e1, e1.Name)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
}

func TestStructOperations(t *testing.T) {
	e:=Employee{"0", "Bobo", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String("传递"))
	t.Log(e.String2("指针"))
}

// 1.这个方法调用实例时,这个实例会复制一份出来用于调用,就浪费内存了
func (e Employee) String(op string) string {
	fmt.Printf("%s address is %x\n", op, unsafe.Pointer(&e.Name))	//使用非指针方式调用实例时,指针位置是不一样的, 说明数据被拷贝了一份
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

// 2.这个方法用实例时，我们直接使用指针来读取数据，就避免了内存浪费
func (e *Employee) String2(op string) string {
	fmt.Printf("%s address is %x\n", op, unsafe.Pointer(&e.Name))	//使用指针方式调用实例时,指针位置是一样的, 使用指针访问的话，可以避免拷贝带来的内存开销
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}