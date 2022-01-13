package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	modulus := make([]int, 0)
	modulusMap := make(map[int]int)
	for i := 0; i < n; i++ {
		var a, b int
		_, _ = fmt.Scanf("(%d,%d)", &a, &b)
		if _, ok := modulusMap[b]; ok == false {
			modulusMap[b] = a
			modulus = append(modulus, b)
		}
	}
	sort.Slice(modulus, func(i, j int) bool {
		return modulus[i] < modulus[j]
	})
	fmt.Println(mPrint(modulus, modulusMap))
}
func mPrint(modulus []int, modulusMap map[int]int) (s string) {
	for i := range modulus {
		if i != 0 && modulusMap[modulus[i]] > 0 {
			s += "+"
		}
		s += strconv.Itoa(modulusMap[modulus[i]])
		if modulus[i] != 0 {
			s += "X"
			if modulus[i] != 1 {
				s += "^"
				s += strconv.Itoa(modulus[i])
			}
		}
	}
	return
}
