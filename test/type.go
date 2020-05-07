package main

import "fmt"

// type后面跟的是类型
type myInt int		// 自定义类型
					// 自定义类型编写完成后依旧有效
type yourInt=int	// 类型别名:还有byte、rune
					// 类型别名只在你代码编写过程中有效，完成后就没有了

type person struct {
	name,gender	string
	age		int
	//gender	string
	hobby	[]string
}

type xyz struct {
	a,b,c int8
}

type student struct {
	name string
	age  int
}

// 标识符：变量名 函数名 类型名 方法名
// Go语言中如果标识符首字母是大写的，就表示对外部包可见（暴露的，共有的）

// dog 这是一条狗的结构体
type dog struct {
	name string
}

// 结构体是值类型，赋值的时候都是拷贝。
// 构造函数：约定成俗用new开头
// 返回的是结构体还是结构体指针
// 当结构体比较大的时候尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) *person {
	return &person{
		name:name,
		age:age,
	}
}

// 方法是作用于特定类型的函数
// 接收者(d)表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog)wang() {
	fmt.Printf("%s:汪汪汪~\n", d.name)
}
func newDog(name string) dog {
	return dog{
		name:name,
	}
}
func main() {
	var n myInt
	n = 100
	fmt.Printf("%T %v\n",n,n)
	var m yourInt
	m = 100
	fmt.Printf("%T %v\n",m,m)
	var c rune
	c = '中'
	fmt.Printf("%T %v\n",c,c)

	// 声明一个person类型的变量p
	var p person
	// 通过字段赋值
	p.name = "zhl"
	p.age = 20
	p.gender = "男"
	p.hobby = []string{"羽毛球", "轮滑"}
	fmt.Printf("type:%T value:%v name:%v\n", p, p, p.name)	//main.person
	var p2 person
	p2.name = "mumu"
	p2.gender = "女"
	fmt.Printf("type:%T value:%v\n", p2, p2)

	// 匿名结构体：多用于临时场景
	var s struct{
		x	string
		y	int
	}
	s.x = "嘿嘿嘿"
	s.y = 100
	fmt.Printf("type:%T value:%v\n", s, s)

	var p3 person
	p3.name = "benben"
	p3.gender = "女"
	f(p3)	//拷贝了一份
	fmt.Printf("真正数据 指针：%p value:%v\n", &p3, p3)
	f2(&p3)
	fmt.Printf("真正数据 指针：%p value:%v\n", &p3, p3)
	// 结构体指针1
	var p4 = new(person)
	p4.name = "yoyo"
	fmt.Printf("type:%T p2的值：%p p2的内存地址：%p\n", p4, p4, &p4)	//值类型的指针 *main.person
	// 结构体指针2
	/*
	取结构体的地址实例化
	使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	*/
	// key-value初始化
	var p5 = &person{	// 声明变量并初始化
		name: "toto",
		gender: "保密",
	}
	fmt.Printf("%#v\n", p5)
	// 使用值列表的形式初始化，值的顺序要和结构体定义时字段的顺序一样
	/*
	使用这种格式初始化时，需要注意：
	1、必须初始化结构体的所有字段。
	2、初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	3、该方式不能和键值初始化方式混用。
	*/
	p6 := &person{
	"momo",
	"女",
	16,
	[]string{"足球", "滑雪"},
	}
	fmt.Printf("%#v\n", p6)

	// 结构体占用一块连续的内存。
	mn := xyz{
		a:int8(10),
		b:int8(20),
		c:int8(30),
	}
	fmt.Printf("a:%p b:%p c:%p\n", &(mn.a), &(mn.b), &(mn.c))

	p1 := newPerson("zhl", 25)
	d1 := newDog("旺财")
	fmt.Println(p1, d1)
	d1.wang()


	m0 := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		fmt.Printf("%p\n", &stu)		// ??????
		m0[stu.name] = &stu
	}
	fmt.Printf("%v\n", m0)
	for k, v := range m0 {
		fmt.Println(k, "=>", v.name)
	}
}
// go语言中函数参数永远是拷贝
func f(x person) {
	x.gender = "男"		// 修改的是副本的gender
	fmt.Printf("副本数据 指针：%p value:%v\n", &x, x)
}
func f2(x *person) {
	//(*x).gender = "男"	// 根据内存地址找到那个原变量，修改的就是原来的变量
	x.gender = "男"		// 语法糖，自动根据指针找对应的变量；go语言中是不能对指针操作的，可以这么写
	fmt.Printf("指针数据 指针：%p value:%v\n", &x, x)	// 这个是一个指向p3指针的指针（np3 = &p3;&np3）
}