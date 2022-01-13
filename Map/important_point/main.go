package main

import (
	"fmt"
	"math"
)

type VexType int

const (
	MaxNum    = 10002
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

func (m *AdjMatrix) AddPoint(x, y int) {
	m.m[x][y] = 1
	m.m[y][x] = 1
}

func (m *AdjMatrix) Floyd() bool {
	if m.edgeNum < m.vexNum-1 {
		return false
	}
	for i := 1; i <= m.vexNum; i++ {
		for j := 1; j <= m.vexNum; j++ {
			for k := i; k <= m.vexNum; k++ {
				if m.m[j][k] > m.m[j][i]+m.m[i][k] {
					m.m[j][k] = m.m[j][i] + m.m[i][k]
				}
				if m.m[k][j] > m.m[k][i]+m.m[i][j] {
					m.m[k][j] = m.m[k][i] + m.m[i][j]
				}
			}
		}
	}
	for i := 1; i <= m.vexNum; i++ {
		for j := 1; j <= m.vexNum; j++ {
			if m.m[i][j] == MaxLength {
				return false
			}
		}
	}
	return true
}

func (m *AdjMatrix) BFS(start, end int) (ret int, ok bool) {
	visited := make([]bool, m.vexNum+1)
	q := []struct{ index, step int }{{start, 0}}
	visited[start] = true
	for len(q) > 0 {
		p := q[0]
		if p.index == end {
			return p.step, true
		}
		q = q[1:]
		for i := 1; i <= m.vexNum; i++ {
			if visited[i] == false && m.m[p.index][i] == 1 {
				visited[i] = true
				q = append(q, struct{ index, step int }{index: i, step: p.step + 1})
			}
		}
	}
	return 0, false
}

func (m *AdjMatrix) Count(x int) (ret int) {
	for i := 1; i <= m.vexNum; i++ {
		ret += m.m[x][i]
	}
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
	flag := adjMatrix.Floyd()
	var num int
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		var vex int
		fmt.Scan(&vex)
		ret := adjMatrix.Count(vex)
		result := float64(adjMatrix.vexNum-1) / float64(ret)
		if flag == false {
			result = 0
		}
		fmt.Printf(Format, vex, result)
	}
}
