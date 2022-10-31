/*
 * @Descripttion :
 * @version      :
 * @Author       :
 * @Date         : 2022-10-26 15:10:49
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2022-10-31 13:25:41
 */
package main

import (
	"fmt"
	"strconv"
)

var num1 int = 99
var num2 float64 = 23.456
var b bool = true
var myChar byte = 'h'
var str string // 空的str

/*
 * fmt.Sprintf() 根据format参数生成格式化的字符串 并返回该字符串
 *
 */
func main() {
	// 基本数据类型转换成string
	// string1()

	// 方式二：strconv 函数
	// string2()

	// string类型转基本数据类型
	// parse()

	//  []byte和字符串互转
	byteString()

}

// 使用第一种方式来转换 fmt.Sprintf方法
func string1() {
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str=%q\n", str, str)
}

// 第二种方式 strconv 函数
func string2() {
	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str type %T str=%q\n", str, str)

	//	f - 格式， 10-表示小数位数保留10位，64-表示这个小数是float64
	str = strconv.FormatFloat(num2, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(b)
	fmt.Printf("str type %T str=%q\n", str, str)

	// 整数转字符串 strconv.Itoa
	str = strconv.Itoa(int(4567))
	fmt.Printf("str type %T str=%q\n", str, str)

	// 字符串转整数 strconv.Atoi
	fmt.Println(strconv.Atoi("12"))
}

func parse() {
	// 字符串转bool
	var s1 string = "false"
	b, _ = strconv.ParseBool(s1)
	fmt.Printf("b type %T str=%q\n", b, b)

	// 字符串转整数
	var s2 string = "12345690"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(s2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type %T n1=%v\n", n1, n1)
	fmt.Printf("n2 type %T n2=%v\n", n2, n2)

	var s3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(s3, 64)
	fmt.Printf("f1 type %T f1=%v\n", f1, f1)
}

func byteString() {
	// []byte转字符串
	fmt.Printf("[]byte 转字符串：%v \n", string([]byte{97, 98, 99}))

	// 字符串遍历，同时处理有中文的问题 r :=[]rune(str)
	r := []rune("hello 上海") // 字符串转[]byte
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}
}
