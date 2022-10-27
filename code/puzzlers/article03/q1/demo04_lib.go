/*
 * 库源码文件是不能被直接运行的源码文件，它仅用于存放程序实体，这些程序实体可以被其他代码使用。
 *
 * 程序实体：变量、常量、函数、结构体和接口的统称。
 */
package main

import "fmt"

func hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
