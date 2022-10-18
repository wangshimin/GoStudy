package main

/****** 通过反射“动态”调用方法 *******/
import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) SayHi(name string) {
	fmt.Printf("Hello,%v! My name is %v.", name, u.Name)
}

func main() {
	u := User{18276, "Tina", 20}
	// u.SayHi("bob")
	v := reflect.ValueOf(u)
	mv := v.MethodByName("SayHi")

	args := []reflect.Value{reflect.ValueOf("joe")}
	mv.Call(args)
}
