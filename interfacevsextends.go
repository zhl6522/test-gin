package main

import "fmt"

type Monkey struct {
	Name	string
}

type LittleMonkey struct {
	Monkey
}

type BridAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (ltm *LittleMonkey) Flying() {
	fmt.Println(ltm.Name, "通过学习，会了飞翔")
}

func (ltm *LittleMonkey) Swimming() {
	fmt.Println(ltm.Name, "通过学习，会了游泳")
}

func (ltm *LittleMonkey) Play() {
	fmt.Println(ltm.Name, "是一只小猴子，在玩耍...")
}

//实现接口 VS 继承
//接口和继承解决的问题不同
//继承的截止主要在于：解决代码的复用性和可维护性。
//接口的截止主要在于：设计，设计好各种规范(方法)，让其他自定义类型去实现这些方法。
//接口比继承更加灵活
//接口比继承更加灵活，继承是满足is - a的关系，而接口只需满足like - a的关系。
//接口在一定程度上实现代码解耦。
func main() {
	monkey := LittleMonkey{Monkey{Name:"悟空"}}
	monkey.Play()
	monkey.Flying()
	monkey.Swimming()
}
