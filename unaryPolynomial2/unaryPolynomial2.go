package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

type num struct {
	modulusMap map[int]int
	modulus    []int
}

type node struct {
	a, b int
}

func main() {
	n1 := input()
	//x := inputX()
	result1 := create(n1)
	fmt.Println(mPrint(derivation(result1)))
}

func create(nodes []node) *num {
	x := &num{
		modulusMap: make(map[int]int),
		modulus:    make([]int, 0),
	}
	for i := range nodes {
		x.modulusMap[nodes[i].b] = nodes[i].a
		x.modulus = append(x.modulus, nodes[i].b)
	}
	return x
}

func input() []node {
	var n int
	_, _ = fmt.Scan(&n)
	//var b byte
	//fmt.Scanf("%c", &b)
	nodes := make([]node, 0, n)
	for i := 0; i < n; i++ {
		var a, b int
		_, _ = fmt.Scanf("(%d,%d)", &a, &b)
		nodes = append(nodes, node{
			a: a,
			b: b,
		})
	}
	//fmt.Scanf("%c", &b)
	return nodes
}

func mPrint(result *num) (s string) {
	sort.Slice(result.modulus, func(i, j int) bool {
		return result.modulus[i] < result.modulus[j]
	})
	for i := range result.modulus {
		if result.modulusMap[result.modulus[i]] == 0 {
			continue
		}
		if i != 0 && result.modulusMap[result.modulus[i]] > 0 {
			s += "+"
		}
		s += strconv.Itoa(result.modulusMap[result.modulus[i]])
		if result.modulus[i] != 0 {
			s += "X"
			if result.modulus[i] != 1 {
				s += "^"
				s += strconv.Itoa(result.modulus[i])
			}
		}
	}
	return
}

func add(result1, result2 *num) *num {
	result := result1
	for _, p := range result2.modulus {
		_, ok := result1.modulusMap[p]
		if ok == false {
			result.modulus = append(result.modulus, p)
		}
		result.modulusMap[p] += result2.modulusMap[p]
	}
	return result
}

func sub(result1, result2 *num) *num {
	result := result1
	for _, p := range result2.modulus {
		_, ok := result1.modulusMap[p]
		if ok == false {
			result.modulus = append(result.modulus, p)
		}
		result.modulusMap[p] -= result2.modulusMap[p]
	}
	return result
}

func mult(result1, result2 *num) *num {
	resultNum := num{
		modulusMap: make(map[int]int),
		modulus:    []int{},
	}
	for _, p := range result2.modulus {
		for _, q := range result1.modulus {
			_, ok := resultNum.modulusMap[p+q]
			if ok == false {
				resultNum.modulus = append(resultNum.modulus, p+q)
			}
			resultNum.modulusMap[p+q] += result2.modulusMap[p] * result1.modulusMap[q]
		}
	}
	return &resultNum
}

func inputX() int {
	var n int
	fmt.Scan(&n)
	//var b int
	//fmt.Scanf("%c", &b)
	return n
}

func eval(n *num, x int) (sum int) {
	for _, p := range n.modulus {
		sum += n.modulusMap[p] * int(math.Pow(float64(x), float64(p)))
	}
	return
}

func derivation(n *num) *num {
	resultNum := num{
		modulusMap: make(map[int]int),
		modulus:    []int{},
	}
	for _, p := range n.modulus {
		if p == 0 {
			continue
		} else {
			resultNum.modulus = append(resultNum.modulus, p-1)
			resultNum.modulusMap[p-1] = n.modulusMap[p] * p
		}
	}
	return &resultNum
}
