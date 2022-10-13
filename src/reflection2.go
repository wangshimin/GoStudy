package main

/****** 反射修改对象的值 *******/
import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	u := User{20256, "lucy", 16}
	set(&u)
	fmt.Println(u)
}

func set(o interface{}) {
	v := reflect.ValueOf(o)

	//	指针类型的接口 && 可被修改
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("XXX")
		return
	} else {
		// 赋值
		v = v.Elem()
	}

	// 找到字段名，修改值
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("没有找到该字段")
		return
	}
	if f.Kind() == reflect.String {
		f.SetString("tara")
	}
}
