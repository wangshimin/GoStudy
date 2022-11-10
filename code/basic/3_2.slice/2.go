/*
 * @Descripttion :
 * @version      :
 * @Author       :
 * @Date         : 2022-11-08 19:16:42
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2022-11-09 16:18:15
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

func example() {
	// string底层是一个byte数组，因此string也可以进行切片处理
	str := "hello@world"
	s := str[6:]
	fmt.Println("slice=", s)
	// 如果需要修改字符串，可以先将string 转换为 ([]byte / []rune)进行修改，再通过string转化为字符串
	arr1 := []byte(s)
	arr1[0] = 'z'
	s = string(arr1)
	fmt.Println("str=", s)

	// 细节，我们转成[]byte后，可以处理英文和数字，但是不能处理中文
	// 原因是 []byte 字节来处理，而一个汉字，是3个字节，因此就会出现乱码
	// 解决方法：将string转成[]rune 即可，因为[]rune是按字符处理，兼容汉字
	arr2 := []rune(str)
	fmt.Print(arr2)
	arr2[0] = '哈'
	s = string(arr2)
	fmt.Println("str=", s)

}

func main() {
	splitString()
	example()
}
