package main

import (
	"fmt"
)

func main() {
	A()
	B()
	C()
}

func A() {
	fmt.Println("FUNC A")
}

func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("wwwww")
		}
	}()
	panic("panic in B")
}

func C() {
	fmt.Println("FUNC C")
}
