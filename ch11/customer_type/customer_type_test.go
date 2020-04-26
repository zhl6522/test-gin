package customer_type

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

// 定义一个函数timeSpent帮助记录入参函数的执行耗时
// 入参是一个函数, 返回也是一个函数
func timeSpent(inner IntConv) IntConv {
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

func TestFun(t *testing.T) {
	// 给timeSpent入参一个函数, 我们返回一个函数, 这个函数执行这个入参函数并打印该的执行时间
	// tsSF 得到的就是返回的函数，我们给这个函数提供一个 「返回函数的参数」-> 10
	tsSF:=timeSpent(slowFun)
	// 打印执行
	t.Log(tsSF(10))
	//t.Log(timeSpent(slowFun)(10))
}