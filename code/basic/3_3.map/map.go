/*
 * @Descripttion :
 * @version      :
 * @Author       :
 * @Date         : 2022-10-12 11:58:27
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2022-11-09 14:01:47
 */
package main

import (
	"fmt"
	"sort"
)

// 语法 var map1 map[keyType]valueType
func main() {
	m := make(map[int]string)
	m[1] = "ok"
	m[8] = "yes"
	delete(m, 1)
	fmt.Println(m, m[1], m[0])

	// sm := make([]map[int]string, 5)
	// for i := range sm {
	// 	sm[i] = make(map[int]string, 1)
	// 	sm[i][2] = "ok"
	// 	fmt.Println(sm[i])
	// }
	// fmt.Println(sm)

	m1 := map[int]string{5: "a", 2: "b", 4: "c", 3: "d", 1: "e"}
	fmt.Println(m1)

	m2 := make(map[string]int)

	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)

	s := make([]int, len(m1))
	i := 0
	for kk := range m1 {
		s[i] = kk
		i++
	}
	sort.Ints(s)
	fmt.Println(s)

	// 判断键值对是否存在
	m3 := map[string]int{"Shanghai": 10, "Beijing": 20}
	value, isset := m3["Shanghai"]
	fmt.Println(value, isset)
	if _, ok := m3["Guangzhou"]; !ok {
		m3["Guangzhou"] = 30
	}
	fmt.Println(m3)

}
