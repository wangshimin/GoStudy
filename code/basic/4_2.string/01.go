/*
 * @Descripttion :
 * @version      :
 * @Author       :
 * @Date         : 2022-10-12 11:58:27
 * @LastEditors  :
 * @LastEditTime : 2022-10-31 13:06:10
 */
/*
string 字符串
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

func exampleFunction() {
	var str string = "Hello,world!"

	// HasPrefix() \ HasSuffix()
	fmt.Printf("在字符串\"Hello,world\"中,\n 以\"He\"开头？ %t\n 以\"He\"结尾？ %t \n", strings.HasPrefix(str, "He"), strings.HasSuffix(str, "He"))

	// Contains()
	fmt.Printf("包含了\"llo\"？ %t \n", strings.Contains(str, "llo"))

	// Index() / LastIndex() / Count()
	fmt.Printf("o最先出现的位置是 %d ，最后出现的位置是 %d ，总共出现了 %d 次\n", strings.Index(str, "o"), strings.LastIndex(str, "o"), strings.Count(str, "o"))

	// Replace()
	fmt.Printf("将字符串中的前2个\"l\"换成\"t\"后：%s \n", strings.Replace(str, "l", "t", 2))

	// ToUpper() / ToLower()
	fmt.Printf("全部转换为大写后：%s \n 全部转换为小写后： %s \n", strings.ToUpper(str), strings.ToLower(str))

	// Trim() / TrimLeft() / TrimRight()
	fmt.Printf("把结尾的\"!\"去掉后：%s \n", strings.Trim(str, "!"))

	// Split() / Fields()
	sp := strings.Split(str, ",")
	fmt.Printf("用,分割字符串，返回了一个slice，索引0:%v，索引1:%v \n", sp[0], sp[1])

	// Join()
	sj := strings.Join(sp, " ## ")
	fmt.Printf("拼接后：%v \n", sj)

	fmt.Printf("长度为：%d 字节\n", len(str))
}

func exampleFormat() {
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
