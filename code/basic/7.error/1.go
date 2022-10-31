/*
 * @Descripttion :
 * @version      :
 * @Author       :
 * @Date         : 2022-10-28 12:26:03
 * @LastEditors  :
 * @LastEditTime : 2022-10-31 14:13:42
 */
package main

import (
	"fmt"
)

/**
 * @Descripttion:  defer + recover来处理错误
 * @return {*}
 */
func exapmle2() {
	defer func() {
		err := recover() //	recover()内置函数，可以捕获到异常
		if err != nil {  // 	说明捕获到异常
			fmt.Println("error=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func main() {
	exapmle2()
}
