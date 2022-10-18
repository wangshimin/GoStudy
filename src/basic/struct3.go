package main

import "fmt"

type TZ int

func (tz *TZ) Increase(num int) {
	*tz += TZ(num) // int型的num强制转换为TZ型，同类型的才能进行运算
}

func main() {
	var a TZ
	a.Increase(100)
	fmt.Println(a)
}
