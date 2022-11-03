/*
 * @Descripttion :
 * @version      :
 * @Author       :
 * @Date         : 2022-10-12 11:58:27
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2022-11-02 16:45:22
 */
package main

import (
	"fmt"
)

func main() {
	A()
	B()
	C()
}

func A() {
	fmt.Println("FUNC A")
}

func B() {
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		if err := recover(); err != nil {
			fmt.Println("wwwww")
		}
	}()
	// panic()触发宕机时，后面的代码不会被运行
	// 前面已经运行过的defer()语句依然会在宕机发生时有效
	panic("panic in B")
}

func C() {
	fmt.Println("FUNC C")
}
