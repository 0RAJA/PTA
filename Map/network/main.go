package main

import (
	"fmt"
	"math"
)

type VexType int

const (
	MaxNum    = 120
	MaxLength = math.MaxInt32
)

type AdjMatrix struct {
	vexs            [MaxNum]VexType
	m               [MaxNum][MaxNum]int
	vexNum, edgeNum int
}

func NewAdjMatrix(vexNum int, edgeNum int) *AdjMatrix {
	M := &AdjMatrix{vexNum: vexNum, edgeNum: edgeNum}
	for i := 1; i <= vexNum; i++ {
		M.vexs[i] = VexType(i)
	}
	for i := 1; i <= vexNum; i++ {
		M.m[i][i] = 1
	}
	return M
}

func (m *AdjMatrix) AddPoint(x, y int) {
	m.m[x][y] = 1
	m.m[y][x] = 1
}

func (m *AdjMatrix) DFS(start, end int) (ret int) {
	visited := make([]bool, m.vexNum)
	visited[start-1] = true
	visited[end-1] = true
	var dfs func(index, cnt int)
	dfs = func(index, cnt int) {
		if cnt == 2 {
			if m.m[index][end] == 1 {
				ret++
			}
			return
		}
		for i := 1; i <= m.vexNum; i++ {
			if m.m[index][i] == 1 && visited[i-1] == false {
				visited[i-1] = true
				dfs(i, cnt+1)
				visited[i-1] = false
			}
		}
	}
	dfs(start, 0)
	return
}

func main() {
	var vexNum, edgeNum int
	fmt.Scanf("%d %d\n", &vexNum, &edgeNum)
	adjMatrix := NewAdjMatrix(vexNum, edgeNum)
	for i := 0; i < edgeNum; i++ {
		var x, y int
		fmt.Scanf("%d %d\n", &x, &y)
		adjMatrix.AddPoint(x, y)
	}
	cnt := 0
	for i := 1; i <= adjMatrix.vexNum; i++ {
		for j := 1; j <= adjMatrix.vexNum; j++ {
			cnt += adjMatrix.DFS(i, j)
		}
	}
	fmt.Println(cnt)
}
