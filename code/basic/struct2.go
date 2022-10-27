package main

import "fmt"

const (
	MAN  int = 1
	LADY int = 2
)

type human struct {
	Sex  int
	Name string
	Age  int
}

type teacher struct {
	//	嵌入结构作为匿名字段,GO会默认将结构名称作为字段名称
	human
	course string
}

type student struct {
	human
}

func main() {
	ta := teacher{course: "math", human: human{Sex: LADY, Name: "lucy"}}
	ta.Age = 18
	// ta.human.Age = 10
	ta.Sex = MAN
	fmt.Println(ta)
}
