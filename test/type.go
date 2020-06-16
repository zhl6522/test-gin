package main

import "fmt"

// type后面跟的是类型
type myInt int		// 自定义类型
					// 自定义类型编写完成后依旧有效
type yourInt=int	// 类型别名，比如:内置的byte、rune
					// 类型别名只在你代码编写过程中有效，完成后就不存在了

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
	fmt.Printf("%p\n", &d)
	fmt.Printf("%s:汪汪汪~\n", d.name)
}
// 使用值接收者：传拷贝进去
func (p person) guonian() {
	p.age++		// 此处p是p1的副本，改的是副本
}
// 指针接收者：传内存地址进去
/*
什么时候应该使用指针类型接收者
1、需要修改接收者中的值
2、接收者是拷贝代价比较大的大对象
3、保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
*/
func (p *person) zhenguonian() {
	p.age++
}
func (p *person) dream(str string) {
	fmt.Printf("%v的梦想：%s\n", p.name, str)
}
func newDog(name string) dog {
	return dog{
		name:name,
	}
}

// 给自定义类型加方法
// 不能给别的包里面的类型添加方法，只能给自己包里的类型添加方法
type myInt10 int

func (m myInt10) helllo() {
	fmt.Println("我是一个int")
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

	fmt.Println("---------------匿名结构体---------------")
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

	fmt.Println("---------------结构体指针---------------")
	// 结构体指针1
	var p4 = new(person)
	p4.name = "yoyo"
	fmt.Printf("type:%T p4的值：%v p4的内存地址：%p\n", p4, p4, &p4)	//值类型的指针 *main.person
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

	fmt.Println("---------------方法---------------")
	p1 := newPerson("zhl", 25)
	d1 := newDog("旺财")	//构造函数
	fmt.Println(p1, d1)
	d1.wang()
	//p1.wang()		// 定义了wang只能Dog接收者才能调用
	p1.guonian()
	fmt.Printf("值传递%v的年龄：%v\n", p1.name,p1.age)
	p1.zhenguonian()
	fmt.Printf("一年后%v的年龄：%v\n", p1.name,p1.age)
	p1.dream("学好Go语言")

	m10 := myInt10(10)	//	var m10 int32 = 10/ var m10 := int32(10)
	m10.helllo()

	m0 := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	for _, stu := range stus {
		fmt.Printf("%p\n", &stu.name)		// 因为stu只有一个内存地址,stu在for range中是同一个变量
		m0[stu.name] = &stu			// 传的是指针的内存地址
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