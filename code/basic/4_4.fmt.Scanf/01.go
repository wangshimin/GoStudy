package main

import "fmt"

// 要求：可以从控制台接收用户信息，【姓名、年龄、薪水、是否通过考试】
// 声明需要的变量
var name string
var age byte
var sal float32
var isPass bool

func main() {

	// 方式1:使用fmt.Scanln()获取
	// type1()

	// 方式2:使用fmt.Scanf()获取
	type2()

}

// 方式1:使用fmt.Scanln()获取
func type1() {
	fmt.Println("请输入姓名 ")
	// 当程序执行到 fmt.Scanln(&name)，程序会停止在这里，等待用户输入，冰回车
	fmt.Scanln(&name)
	fmt.Println("请输入年龄 ")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水 ")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过考试 ")
	fmt.Scanln(&isPass)
	fmt.Printf("名字：%v \n 年龄：%v\n 薪水：¥%v \n 是否通过考试：%v \n", name, age, sal, isPass)
}

// 方式2: fmt.Scanf，可以按指定的格式输入
func type2() {
	fmt.Println("请输入您的姓名，年龄，薪水，是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("名字：%v \n 年龄：%v\n 薪水：¥%v \n 是否通过考试：%v \n", name, age, sal, isPass)
}
