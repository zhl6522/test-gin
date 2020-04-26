package fun__test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10),rand.Intn(20)
}

func TestFun(t *testing.T) {
	a,b:=returnMultiValues()
	t.Log(a, b)
	// 给timeSpent入参一个函数, 我们返回一个函数, 这个函数执行这个入参函数并打印该的执行时间
	// tsSF 得到的就是返回的函数，我们给这个函数提供一个 「返回函数的参数」-> 10
	tsSF:=timeSpent(slowFun)
	// 打印执行
	t.Log(tsSF(10))
	//t.Log(timeSpent(slowFun)(10))
}

// 定义一个函数timeSpent帮助记录入参函数的执行耗时
// 入参是一个函数, 返回也是一个函数
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start:=time.Now()
		fmt.Println("time start:", start)
		ret:=inner(n)		//ret := 入参函数名(返回函数的参数) // 执行有入参的该函数
		fmt.Println("该入参函数执行结果：", ret)
		fmt.Println("时间损耗:", time.Since(start).Seconds())	// 打印计算出的 函数执行花费的时间
		end := time.Now()
		fmt.Println("time end:", end)
		return ret
	}
}

// 延时一秒后返回入参
func slowFun(op int) int {
	time.Sleep(time.Second*1)
	return op
}

func Sum(ops ...int) int {
	ret := 0
	for _,v := range ops {
		ret += v
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1,2,3,4))
	t.Log(Sum(1,2,3,4,5))
}

func Clear() {
	fmt.Println("Clear resources.")
}

// defer「延缓」碰到这个函数，先延缓不动。在该函数最后返回前，才执行defer声明的函数, 一般用来释放执行函数后的资源和锁
// panic「惊恐」让人感到惊恐的错误 ⚠️：即便出现严重错误，defer 依然可以执行，用于释放资源
func TestDefer(t *testing.T) {
	defer Clear()		//延迟执行
	fmt.Println("Start")
	//panic("err")
}

func Clear2(par1 int) {
	fmt.Println("3. Clear resources!", par1)
}

func FuncDefer() {
	fmt.Println("2. 被 defer 的函数里面还有被 defer 的化，会被先执行!")
	defer Clear2(2)
}

func TestDefer2(t *testing.T) {
	defer Clear2(1)
	defer FuncDefer()
	fmt.Println("1. Start")
	panic("err")
}
