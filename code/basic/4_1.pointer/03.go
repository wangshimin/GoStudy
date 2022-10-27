/* 示例：创建指针的另一种方法：new()函数
*
* new()函数可以创建一个对应类型的指针，创建过程会分配内存。
* 被创建的指针指向的值为默认值。
 */
package main

import "fmt"

func main() {
	str := new(string)
	*str = "ninja"
	fmt.Println(*str)
}
