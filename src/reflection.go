package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Mannager struct {
	User  //	匿名字段
	title string
}

func (u User) Hello() {
	fmt.Println("Hello world.")
}

func main() {
	person := User{77, "bob", 18}
	Info(person)

	m := Mannager{User: User{22, "lucy", 16}, title: "总经理"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v \n %#v \n", t.Field(0), t.Field(1))
	fmt.Printf("%#v \n", t.FieldByIndex([]int{0, 1})) // 获取匿名字段里的信息

	xp := 123
	xpv := reflect.ValueOf(&xp)
	xpv.Elem().SetInt(999) // 指针反射：Elem()
	fmt.Println(xp)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o) // 通过反射的typeof获取接收者的类型
	fmt.Println("Type:", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields:", t.Kind() == reflect.Struct)

	//	获取对象的字段、类型、值
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v \n", f.Name, f.Type, val)
	}

	// 获取对象的方法名、方法签名
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v \n", m.Name, m.Type)
	}

}
