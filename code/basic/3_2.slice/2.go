/*
 * @Descripttion :
 * @version      :
 * @Author       :
 * @Date         : 2022-11-08 19:16:42
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2022-11-08 19:26:18
 */
package main

import "fmt"

/** 例子：编写一个函数，要求其接受两个参数，原始字符串 str 和分割索引 i，然后返回两个分割后的字符串。
 * @Descripttion:
 * @return {*}
 */
func splitString() {
	var Split = func(s string, pos int) (string, string) {
		return s[0:pos], s[pos:]
	}
	str := "Google"
	for i := 0; i < len(str); i++ {
		a, b := Split(str, i)
		fmt.Printf("The string %s split at position %d is: %s / %s\n", str, i, a, b)
	}
}

func main() {
	splitString()

}
