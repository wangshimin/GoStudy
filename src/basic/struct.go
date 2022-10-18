package main

import "fmt"

type person struct {
	Name string
	Age  int
	//	匿名结构
	Contact struct {
		Phone, City string
	}
}

type man struct {
	//	结构里的匿名字段
	string
	int
}

func main() {
	a := &person{
		Name: "Tara",
		Age:  18,
	}
	a.Name = "OK"
	a.Age = 19
	a.Contact.Phone = "18221881188"
	a.Contact.City = "ShangHai"

	// fmt.Println(a)
	// A(a)
	// B(a)
	// fmt.Println(a)

	//	匿名结构
	anonymousA := &struct {
		Name string
		Age  int
	}{
		Name: "Bob",
		Age:  20,
	}
	anonymousA.Name = "lucy"
	// fmt.Println(anonymousA)

	//	支持匿名字段
	c := man{"PP", 55}
	fmt.Println(c)
}

func A(per *person) {
	per.Age = 13
	fmt.Println("A", per)
}

func B(per *person) {
	per.Age = 15
	fmt.Println("B", per)
}
