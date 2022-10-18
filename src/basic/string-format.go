package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	/************ 格式化字符串 *****************/
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", true)
	fmt.Printf("%d\n", 123)
	fmt.Printf("%b\n", 14)
	fmt.Printf("%c\n", 96)
	fmt.Printf("%x\n", 456)
	fmt.Printf("%X\n", 456)
	fmt.Printf("%f\n", 456.8889900)
	fmt.Printf("%e\n", 456.8889900)
	fmt.Printf("%E\n", 456.8889900)
	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%x\n", "hey this")
	fmt.Printf("%p\n", &p)
	fmt.Printf("|%6d|%6d|\n", 345, 13)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
