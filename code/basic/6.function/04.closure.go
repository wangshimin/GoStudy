/*
 * @Descripttion : 闭包
 * @version      :
 * @Author       :
 * @Date         : 2022-10-27 16:27:42
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2022-10-31 12:07:00
 */
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

// 示例：闭包的记忆效应
// 被捕获到的闭包中的变量让闭包本身拥有了记忆效应
// 闭包中的逻辑可以修改闭包捕获的变量
// 变量会跟随生命期一直存在
// 闭包本身就如同变量一样拥有了记忆效应
func example1() {
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

// 累加器
func AddUpper() func(int) int {
	var n int = 10
	var str = "hello"
	// 返回一个匿名函数，但这个匿名函数引用到函数外的n
	// 因此这个匿名函数就和n形成一个整体，构成闭包
	return func(x int) int {
		n = n + x
		str += string(36)        // 36 = '$'
		fmt.Println("str=", str) // hello$ hello$$ hello$$$
		return n
	}
}

// 创建一个玩家生成器，输入名称，输出生成器
func playerGen(name string) func() (string, int) {
	// 血量一直为150
	hp := 150

	// 返回创建的闭包
	return func() (string, int) {
		// 将变量引用到闭包中
		return name, hp
	}
}

func example2() {
	f := AddUpper()
	fmt.Println(f(1)) //	11
	fmt.Println(f(2)) //	13
	fmt.Println(f(3)) //	16
}

// 示例：闭包实现生成器
func examplePlayerGen() {
	//	创建一个玩家生成器
	gtor := playerGen("bob")

	// 返回玩家的名字和血量
	name, hp := gtor()

	// 打印值
	fmt.Println(name, hp)
}
func main() {
	// example1()
	example2()
	// examplePlayerGen()
}
