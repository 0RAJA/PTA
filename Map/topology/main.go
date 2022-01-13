package main

import "fmt"

type VexType byte

const (
	MaxNum = 100
)

type Stack []int

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
	m.vexs[index1].outDegree++
	m.vexs[index2].inDegree++
}
func (m *AdjMatrix) TopologicalSort() (ret string, ok bool) {
	stack := Stack{}
	for i := 1; i <= m.vexNum; i++ {
		if m.vexs[i].inDegree == 0 {
			stack.Push(i)
		}
	}
	for !stack.IsEmpty() {
		index := stack.Pop()
		ret += string(m.vexs[index].vex)
		for i := 1; i <= m.vexNum; i++ {
			if m.m[index][i] == 1 {
				m.vexs[i].inDegree--
				if m.vexs[i].inDegree == 0 {
					stack.Push(i)
				}
			}
		}
	}
	if len(ret) < m.vexNum {
		return "", false
	}
	return ret, true
}
func main() {
	vexNum, edgeNum := 0, 0
	fmt.Scanf("%d %d\n", &vexNum, &edgeNum)
	var s string
	fmt.Scanln(&s)
	adjMatrix := NewAdjMatrix(vexNum, edgeNum, s)
	for i := 0; i < edgeNum; i++ {
		var x, y VexType
		fmt.Scanf("<%c,%c>", &x, &y)
		adjMatrix.Add(x, y)
	}
	str, _ := adjMatrix.TopologicalSort()
	fmt.Println(str)
}

func (s *Stack) Push(index int) {
	*s = append(*s, index)
}

func (s *Stack) Pop() (ret int) {
	ret = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
