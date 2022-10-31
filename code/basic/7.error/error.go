/*
 * @Descripttion : 自定义错误
 * @version      :
 * @Author       :
 * @Date         : 2022-10-28 12:26:03
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2022-10-31 14:14:57
 */
package main

import (
	"errors"
	"fmt"
)

// 声明一个解析错误
type ParseError struct {
	Filename string //	文件名
	Line     int    //	行号
}

// 实现error接口，返回错误描述
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line)
}

// 创建一个错误实例
func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}

// 定义除数为0的错误
var errDivisionByZero = errors.New("division by zero")

/**
 * @Descripttion:           除法函数，除数为0时，返回一个预定义的除数为0的错误
 * @param {*} dividend		被除数
 * @param {int} divisor		除数
 * @return {int}			商
 * @return {error} 			错误信息
 */
func div(dividend, divisor int) (int, error) {
	// 判断除数为0的情况并返回
	if divisor == 0 {
		return 0, errDivisionByZero
	}
	// 正常计算，返回空错误
	return dividend / divisor, nil
}

func example1() {
	arr := [2][2]int{{10, 2}, {3, 0}}
	for _, v := range arr {
		res, err := div(v[0], v[1])
		if err != nil {
			errorMsg("test.go", 46)
			return
		}
		fmt.Printf("%d/%d，结果为：%d  \n", v[0], v[1], res)
	}
}

/**
 * @Descripttion: 自定义错误信息
 * @param {string} fileName
 * @param {int} line
 * @return {*}
 */
func errorMsg(fileName string, line int) {
	var e error
	//	创建一个错误实例，包含文件名和行号
	e = newParseError(fileName, line)

	// 通过error接口查看错误描述
	fmt.Println(e.Error())

	// 根据错误接口的具体类型，获取详细的错误信息
	switch detail := e.(type) {
	case *ParseError:
		// 这是一个解析错误
		fmt.Printf("Filename: %s Line:%d\n", detail.Filename, detail.Line)
	default:
		fmt.Println("other error")
	}
}

func main() {
	example1()
}
