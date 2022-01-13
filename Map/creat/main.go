package main

import (
	"fmt"
	"math"
)

type VexType byte

const (
	MaxNum    = 25
	MaxLength = math.MaxInt32
	Format    = "%c %d %d\n"
)

type vexNode struct {
	vex                 VexType
	inDegree, outDegree int
}

type AdjMatrix struct {
	vexs            [MaxNum]vexNode
	m               [MaxNum][MaxNum]int
	vexNum, edgeNum int
}

func NewAdjMatrix(vexNum int, edgeNum int, vexs string) *AdjMatrix {
	M := &AdjMatrix{vexNum: vexNum, edgeNum: edgeNum}
	for i := 1; i <= vexNum; i++ {
		M.vexs[i].vex = VexType(vexs[i-1])
	}
	return M
}

func (m *AdjMatrix) Add(x, y VexType) {
	var index1, index2 int
	for i := 1; i <= m.vexNum; i++ {
		if m.vexs[i].vex == x {
			index1 = i
			break
		}
	}
	for i := 1; i <= m.vexNum; i++ {
		if m.vexs[i].vex == y {
			index2 = i
			break
		}
	}
	m.m[index1][index2] = 1
}

func (m *AdjMatrix) Count() {
	for i := 1; i <= m.vexNum; i++ {
		for j := 1; j <= m.vexNum; j++ {
			if m.m[i][j] == 1 {
				m.vexs[i].outDegree++
				m.vexs[j].inDegree++
			}
		}
	}
}
func (m *AdjMatrix) DFS(start int) (ret string) {
	visited := make([]bool, m.vexNum)
	var dfs func(index int)
	dfs = func(index int) {
		visited[index-1] = true
		ret += string(m.vexs[index].vex)
		for i := 1; i <= m.vexNum; i++ {
			if m.m[index][i] == 1 && visited[i-1] == false {
				dfs(i)
			}
		}
	}
	dfs(start)
	return
}

func (m *AdjMatrix) BFS(start int) (ret string) {
	visited := make([]bool, m.vexNum)
	q := []int{start}
	visited[start-1] = true
	for len(q) > 0 {
		p := q[0]
		ret += string(m.vexs[p].vex)
		q = q[1:]
		for i := 1; i <= m.vexNum; i++ {
			if visited[i-1] == false && m.m[p][i] == 1 {
				visited[i-1] = true
				q = append(q, i)
			}
		}
	}
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
	//for i := 1; i <= adjMatrix.vexNum; i++ {
	//	for j := 1; j <= adjMatrix.vexNum; j++ {
	//		fmt.Print(adjMatrix.m[i][j], " ")
	//	}
	//	fmt.Println()
	//}
	adjMatrix.Count()
	for i := 1; i <= adjMatrix.vexNum; i++ {
		fmt.Printf(Format, adjMatrix.vexs[i].vex, adjMatrix.vexs[i].outDegree, adjMatrix.vexs[i].inDegree)
	}
	fmt.Println(adjMatrix.DFS(1))
	fmt.Println(adjMatrix.BFS(1))
}
