/*
 * @Descripttion :结构体实现接口
 * @version      :
 * @Author       :
 * @Date         : 2022-10-27 15:53:36
 * @LastEditors  :
 * @LastEditTime : 2022-10-28 15:16:32
 */
package main

import "fmt"

// 调用器接口
type Invoker interface {
	// 需要实现一个Call()方法
	// 调用时会传入一个interface{}类型的变量
	// 这种类型的变量表示任意类型的值
	Call(interface{})
}

// 结构体类型
type Struct struct {
}

// 实现Invoker的Call
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

func main() {
	// 声明接口变量
	var Ier Invoker

	// 实例化结构体
	s := new(Struct) // 也可以写成 s := &Struct

	// 将实例化的结构体赋值到接口
	Ier = s // s类型为*Struct,已经实现了Invoker接口的类型，因此赋值给Ier时是成功的

	// 使用接口调用实例化结构体的方法Struct.Call
	Ier.Call("hello") // 通过接口的Call()方法，传入hello，此时将调用Struct结构体的Call()方法

}
