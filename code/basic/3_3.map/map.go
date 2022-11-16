/*
 * @Descripttion :
 * @version      :
 * @Author       :
 * @Date         : 2022-10-12 11:58:27
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2022-11-10 14:42:18
 */
package main

import (
	"fmt"
	"sort"
)

// 语法 var map1 map[keyType]valueType
// map的声明是不会分配内存的，初始化需要make，分配内存后才能赋值和使用
func main() {
	// map的声明和注意事项
	var m map[int]string
	// 在使用map前，需要先make，make的作用就是给map分配数据空间
	m = make(map[int]string)
	m[1] = "ok"
	m[8] = "yes"
	// 演示删除
	delete(m, 1)
	// 当delete指定的key不存在时，删除不会操作，也不会报错
	delete(m, 10)
	fmt.Println(m, m[1], m[0])
	// 如果希望一次性删除所有的key
	// 1.遍历所有的key，逐一删除
	// 2.直接make一个新的空间
	m = make(map[int]string)
	fmt.Println(m)

	fmt.Println("----------- 案例二 ----------")
	sm := make([]map[int]string, 5) // 第二种声明方式
	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][2] = "ok"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)

	fmt.Println("----------- 案例三 ----------")
	m1 := map[int]string{5: "a", 2: "b", 4: "c", 3: "d", 1: "e"} // 第三种声明方式
	fmt.Println(m1)

	m2 := make(map[string]int)

	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)

	/** 排序 **/
	s := make([]int, len(m1))
	i := 0
	for kk := range m1 {
		s[i] = kk
		i++
	}
	sort.Ints(s)
	fmt.Println(s)

	fmt.Println("----------- 案例四 ----------")

	// 判断键值对是否存在
	m3 := map[string]int{"Shanghai": 10, "Beijing": 20}
	value, isset := m3["Shanghai"]
	fmt.Println(value, isset)
	if _, ok := m3["Guangzhou"]; !ok {
		m3["Guangzhou"] = 30
	}
	fmt.Println(m3)

}
