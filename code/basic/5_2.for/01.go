package main

import "fmt"

func main() {
	example1()
	example2()
	example3()
	example4()
}

func example1() {
	a := 1

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(a)
	// 	a++
	// }

	for ; a <= 3; a++ {
		fmt.Println(a)
	}

}

// 结束循环时带可执行语句的无限循环
func example2() {
	a := 1
	for ; ; a++ {
		fmt.Println(a)
		if a > 2 {
			break
		}
	}
}

// example2升级版 - 无限循环：在收发处理中较为常见，但需要无限循环有可控的退出方式来结束循环
func example3() {

	a := 1

	for {
		fmt.Println(a)
		if a >= 3 {
			break
		}
		a++
	}
}

// example3升级版 - 只有一个循环条件的循环
func example4() {
	a := 1
	for a <= 3 {
		fmt.Println(a)
		a++
	}
}
