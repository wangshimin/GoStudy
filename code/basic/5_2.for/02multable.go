/*
 * @Descripttion :九九乘法表
 * @version      :
 * @Author       :
 * @Date         : 2022-10-27 12:00:37
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2022-10-31 12:58:37
 */
package main

import "fmt"

/**
 * @Descripttion: For循环实现
 */
func exampleFor() {

	// 遍历，决定处理第几行
	for y := 1; y <= 9; y++ {

		// 遍历，决定这一行有多少列
		for x := 1; x <= y; x++ {
			fmt.Printf("%d x %d = %d   ", x, y, x*y)
		}

		//	手动生成回车
		fmt.Println()
	}
}

/**
 * @Descripttion: 函数封装后实现
 * @return {*}
 */
func exapleFunction() {
	//	从终端输入一个整数表示要打印的乘法表对应的数
	var num int
	fmt.Println("请输入九九乘法表的对应数")
	fmt.Scanln(&num)
	f := func(n int) {
		for i := 1; i < n; i++ {
			for j := 1; j <= i; j++ {
				fmt.Printf("%v x %v = %v  ", j, i, j*i)
			}
			fmt.Println()
		}
	}
	f(num)
}

func main() {
	exampleFor()
	exapleFunction()
}
