package main

import (
	"fmt"
	"sort"
)

func main() {
	/******************* map **************************/
	// m := make(map[int]string)
	// m[1] = "ok"
	// m[8] = "yes"
	// delete(m, 1)
	// fmt.Println(m, m[1], m[0])

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

}
