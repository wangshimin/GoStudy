// 示例：闭包的记忆效应
// 被捕获到的闭包中的变量让闭包本身拥有了记忆效应
// 闭包中的逻辑可以修改闭包捕获的变量
// 变量会跟随生命期一直存在
// 闭包本身就如同变量一样拥有了记忆效应
package main

import "fmt"

// 提供一个值，每次调用函数会指定对值进行累加
func Accumulate(value int) func() int {
	// 返回一个闭包
	return func() int {
		// 累加
		value++

		// 返回一个累加值
		return value
	}
}

func main() {

	// 创建一个累加器，初始值为1
	ator := Accumulate(1)

	// 累加1并打印
	fmt.Println(ator()) // 2

	fmt.Println(ator()) // 3

	// 打印累加器的函数地址
	fmt.Printf("%p\n", ator)

	// 创建一个累加器，初始值为1
	atro2 := Accumulate(10)

	// 累加1并打印
	fmt.Println(atro2()) // 11

	// 打印累加器的函数地址
	fmt.Printf("%p\n", atro2)

}
