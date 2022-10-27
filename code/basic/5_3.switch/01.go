package main

import "fmt"

func main() {
	// example1()
	// example2()
	// example3()
	// example4()
	// example5()
	TypeSwitch()
}

func example1() {
	b := 1
	switch b {
	case 0:
		fmt.Println("b=0")
	case 1:
		fmt.Println("b=1")
	}
}

func example2() {
	b := 1
	switch {
	case b >= 0:
		fmt.Println("b大于等于0")
		fallthrough
	case b >= 1:
		fmt.Println("b大于1")
	}
}

func example3() {
	switch b := 1; {
	case b >= 0:
		fmt.Println("b大于等于0")
		fallthrough
	case b >= 1:
		fmt.Println("b大于1")
	}
}

func example4() {
LABLE1:
	for {
		for i := 1; i < 10; i++ {
			fmt.Println(i)
			if i > 3 {
				break LABLE1
			}
		}
	}
}

func example5() {
	for c := 0; c < 10; c++ {
		fmt.Printf("第%v行\n", c)
		switch c {
		case 1, 3, 5: // 一分支多值
			fmt.Printf("哈哈哈，我现在是 %v !\n", c)
		case 2:
			fmt.Printf("我现在是 %v !", c)
			fallthrough
		case 4:
			fmt.Printf("***%v****\n", c)
		}
	}
}

func TypeSwitch() {
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型：%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}
