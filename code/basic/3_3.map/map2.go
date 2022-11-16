/* map类型切片
 * @Descripttion :
 * @version      :
 * @Author       :
 * @Date         : 2022-11-08 19:48:39
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2022-11-09 13:58:52
 */
package main

import (
	"fmt"
	"strconv"
)

func example1() {
	// 方式1：通过索引使用切片的map元素
	items := make([]map[int]string, 5) // map类型的切片
	for i := range items {             // 遍历给每个切片分配map
		items[i] = make(map[int]string, 1) //	map
		items[i][1] = "hi"
	}
	fmt.Printf("方式1 - items:%v\n", items)

	// 方式2:不推荐这种写法
	items2 := make([]map[int]string, 5)
	for _, v := range items {
		v = make(map[int]string, 1) // v只是获得map值的一个拷贝
		v[1] = "hello"              // map元素没有得到初始化
	}
	fmt.Printf("方式1 - items2:%v\n", items2)
}

/*
切片的数据类型如果是map，则我们称为slice of map，map切片，
这样使用则【 !! map切片个数就可以动态变化 !! 】了
*/
func exampleSliceOfMap() {
	monster := make([]map[string]string, 2) // map类型的切片

	for i := 0; i < 4; i++ {
		fmt.Println("one")
		// 使用append动态增加，并实现了扩容
		monster = append(monster, map[string]string{
			"name": "t" + strconv.Itoa(i),
			"age":  "18",
		})
	}
	fmt.Printf("monster = %#v\n cap： %d ，len：%d", monster, cap(monster), len(monster))
}

func main() {
	example1()
	exampleSliceOfMap()
}
