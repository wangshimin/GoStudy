/*
 * @Descripttion :
 * @version      :
 * @Author       :
 * @Date         : 2022-10-12 11:58:27
 * @LastEditors  :
 * @LastEditTime : 2022-10-31 18:25:27
 */
package main

import (
	"fmt"
)

type TZ int //	定义底层类型为int的TZ类型

type A struct {
	Name string
	age  int // 私有字段
}

type B struct {
	Name string
}

func main() {
	Va := A{}
	Va.Name = "初始值anna"
	Va.Print()
	fmt.Println(Va.Name, Va.age)

	Vb := B{}
	Vb.Name = "初始值bob"
	Vb.Print()
	fmt.Println(Vb.Name)

	var Vt TZ
	Vt.Print()     // Method Value
	(TZ).Print(Vt) // Method Expression

}

// receiver为类型的指针
func (recevierA *A) Print() {
	recevierA.Name = "AA"
	recevierA.age = 20
	fmt.Println("方法A", recevierA)
}

// receiver为类型的值
func (recevierB B) Print() {
	recevierB.Name = "BB"
	fmt.Println("方法B", recevierB)
}

// 给类型TZ定义一个名为Print的方法
func (receverT TZ) Print() {
	fmt.Println("TZ")
}
