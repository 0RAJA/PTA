package main

import (
	"fmt"
	"math"
)

type VexType byte

const (
	MaxNum    = 25
	MaxLength = math.MaxInt32
	Format    = "%d"
)

type AdjMatrix struct {
	vexs            [MaxNum]VexType
	m               [MaxNum][MaxNum]int
	vexNum, edgeNum int
}

func NewAdjMatrix(vexNum int, edgeNum int, vexs string) *AdjMatrix {
	M := &AdjMatrix{vexNum: vexNum, edgeNum: edgeNum}
	for i := 1; i <= vexNum; i++ {
		M.vexs[i] = VexType(vexs[i-1])
	}
	return M
}

func (m *AdjMatrix) Add(x, y VexType) {
	var index1, index2 int
	for i := 1; i <= m.vexNum; i++ {
		if m.vexs[i] == x {
			index1 = i
			break
		}
	}
	for i := 1; i <= m.vexNum; i++ {
		if m.vexs[i] == y {
			index2 = i
			break
		}
	}
	m.m[index1][index2] = 1
	m.m[index2][index1] = 1
}

func (m *AdjMatrix) DFS(start int, visited []bool) (ret string) {
	var dfs func(index int)
	dfs = func(index int) {
		visited[index-1] = true
		ret += string(m.vexs[index])
		for i := 1; i <= m.vexNum; i++ {
			if m.m[index][i] == 1 && visited[i-1] == false {
				dfs(i)
			}
		}
	}
	dfs(start)
	return
}

func main() {
	var vexNum, edgeNum int
	fmt.Scanf("%d %d\n", &vexNum, &edgeNum)
	var vexs string
	fmt.Scanln(&vexs)
	adjMatrix := NewAdjMatrix(vexNum, edgeNum, vexs)
	for i := 0; i < edgeNum; i++ {
		var s string
		fmt.Scanln(&s)
		adjMatrix.Add(VexType(s[0]), VexType(s[1]))
	}
	cnt := 0
	visited := make([]bool, adjMatrix.vexNum)
	for i := 1; i <= adjMatrix.vexNum; i++ {
		if visited[i-1] == false {
			cnt++
			adjMatrix.DFS(i, visited)
		}
	}
	fmt.Println(cnt)
}
