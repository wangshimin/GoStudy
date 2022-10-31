/*
 * @Descripttion : 递归函数
 * @version      :
 * @Author       :
 * @Date         : 2022-10-28 14:43:38
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2022-10-28 15:14:55
 */
package main

import "fmt"

func main() {
	exampleFbn()
}

func exampleFbn() {
	fmt.Println("|  n | Fbn |\n")
	for i := 1; i < 20; i++ {
		fmt.Printf("| %-2d | %-4d |\n", i, fbn(i))
	}
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
