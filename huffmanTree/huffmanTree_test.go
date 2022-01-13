package main

//
//import (
//	"fmt"
//	"strconv"
//	"strings"
//	"testing"
//)
//
//func TestNewHFMTree(t *testing.T) {
//	str := ""
//	fmt.Scanln(&str)
//	weights := strings.Split(str, " ")
//	weight := make(map[rune]int)
//	for i := 'A'; i <= 'F'; i++ {
//		weight[i], _ = strconv.Atoi(weights[i-'A'])
//	}
//	tree := NewHFMTreeWithWright(weight)
//	for i := 'A'; i <= 'F'; i++ {
//		fmt.Printf("%s:%s", string(i), tree.code[i])
//	}
//	fmt.Scanln(&str)
//	fmt.Println(tree.ToCode(str))
//	fmt.Scanln(&str)
//	fmt.Println(tree.DeCode(str))
//}
