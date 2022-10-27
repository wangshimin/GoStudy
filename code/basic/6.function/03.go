// 函数体实现接口
package main

import "fmt"

// 调用器接口
type Invoker interface {
	// 需要实现一个Call()方法
	// 调用时会传入一个interface{}类型的变量
	// 这种类型的变量表示任意类型的值
	Call(interface{})
}

// 函数定义为类型
type FuncCaller func(interface{})

// 实现Invoker的Call
func (f FuncCaller) Call(p interface{}) {
	// 调用f()函数本体
	f(p)
}

func main() {
	// 声明接口变量
	var Ier Invoker

	// 将匿名函数转为FuncCaller类型，再赋值给接口
	Ier = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})

	// 使用接口调用FuncCaller.Call，内部会调用函数本体
	Ier.Call("hello")

}
