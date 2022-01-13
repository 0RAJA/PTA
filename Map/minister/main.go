package main

import (
	"fmt"
	"math"
)

type VexType int

const (
	MaxNum    = 1000
	MaxLength = math.MaxInt32
	Format    = "Cc(%d)=%.2f\n"
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
		for j := 1; j <= vexNum; j++ {
			if i == j {
				continue
			}
			M.m[i][j] = MaxLength
			M.m[j][i] = MaxLength
		}
	}
	return M
}

func (m *AdjMatrix) AddPoint(x, y, z int) {
	m.m[x][y] = z
	m.m[y][x] = z
}

func count(num int) int {
	return num*10 + (1+num)*num/2
}

func (m *AdjMatrix) DFS(start, end int) (ret int) {
	visited := make([]bool, m.vexNum+1)
	visited[start] = true
	var dfs func(index, sum int)
	dfs = func(index, sum int) {
		if index == end {
			if ret < sum {
				ret = sum
			}
			return
		}
		for i := 1; i <= m.vexNum; i++ {
			if m.m[index][i] != MaxLength && visited[i] == false {
				visited[i] = true
				dfs(i, sum+m.m[index][i])
				visited[i] = false
			}
		}
	}
	dfs(start, 0)
	return
}

func main() {
	var vexNum, edgeNum int
	fmt.Scanf("%d\n", &vexNum)
	edgeNum = vexNum - 1
	adjMatrix := NewAdjMatrix(vexNum, edgeNum)
	for i := 0; i < edgeNum; i++ {
		var x, y, z int
		fmt.Scanf("%d %d %d\n", &x, &y, &z)
		adjMatrix.AddPoint(x, y, z)
	}
	ret := 0
	for i := 1; i <= adjMatrix.vexNum; i++ {
		for j := 1; j <= adjMatrix.vexNum; j++ {
			if i == j {
				continue
			}
			if x := adjMatrix.DFS(i, j); x > ret {
				ret = x
			}
		}
	}
	fmt.Println(count(ret))
}
