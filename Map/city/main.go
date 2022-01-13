package main

import (
	"fmt"
	"sort"
)

type EdgeNode struct {
	begin, end int
	Weight     int
}

type EdgeArray struct {
	list            []EdgeNode
	vexNum, edgeNum int
}

func NewEdgeArray(vexNum int, edgeNum int) *EdgeArray {
	return &EdgeArray{vexNum: vexNum, edgeNum: edgeNum}
}

func (array *EdgeArray) Add(x, y, z int) {
	array.list = append(array.list, EdgeNode{
		begin:  x,
		end:    y,
		Weight: z,
	})
}
func find(parent []int, num int) int {
	if parent[num] == num {
		return num
	} else {
		parent[num] = find(parent, parent[num])
		return parent[num]
	}
}
func (array *EdgeArray) Kruskal() (ret int) {
	sort.Slice(array.list, func(i, j int) bool {
		return array.list[i].Weight < array.list[j].Weight
	})
	parent := make([]int, array.vexNum+1)
	for i := range parent {
		parent[i] = i
	}
	for i := 0; i < len(array.list); i++ {
		m := find(parent, array.list[i].begin)
		n := find(parent, array.list[i].end)
		if m != n {
			parent[n] = m
			ret += array.list[i].Weight
		}
	}
	return
}

func main() {
	var vexNum, edgeNum int
	fmt.Scanf("%d %d\n", &vexNum, &edgeNum)
	edgeArray := NewEdgeArray(vexNum, edgeNum)
	for i := 0; i < edgeNum; i++ {
		var x, y, z int
		fmt.Scanf("%d %d %d\n", &x, &y, &z)
		edgeArray.Add(x, y, z)
	}
	for i := 1; i <= vexNum; i++ {
		var z int
		fmt.Scan(&z)
		if z == -1 {
			continue
		}
		edgeArray.Add(0, i, z)
	}
	fmt.Println(edgeArray.Kruskal())
}
