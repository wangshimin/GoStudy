/*
 * @Descripttion : 递归函数
 * @version      :
 * @Author       :
 * @Date         : 2022-10-28 14:43:38
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2022-11-02 16:12:58
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	exampleFbn()

	exampleFbnBYClosure()
}

func exampleFbn() {
	start := time.Now()
	fmt.Println("|  n | Fbn |\n")
	for i := 1; i < 20; i++ {
		fmt.Printf("| %-2d | %-4d |\n", i, fbn(i))
	}
	end := time.Now()
	fmt.Printf("耗时：%s\n", end.Sub(start))
}

/**
 * @Descripttion: 递归的方式，求出斐波那契数1,1,2,3,5,8,13...
 * @param {int} n 整数n
 * @return {*}    斐波那契数
 */
func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

func fib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

/**
 * @Descripttion:  闭包的方式实现斐波那契数
 * @return {*}
 */
func exampleFbnBYClosure() {
	start := time.Now()
	f := fib()
	for i := 0; i <= 9; i++ {
		println(i+2, f())
	}
	end := time.Now()
	fmt.Printf("耗时：%s\n", end.Sub(start))

}
