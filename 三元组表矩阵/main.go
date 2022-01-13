package main

import (
	"fmt"
	"sort"
)

type MatrixNode struct {
	row, col, value int
}

type Matrix struct {
	Nodes                    []MatrixNode
	MaxRows, MaxCols, Length int
}

func NewMatrixNodes() []MatrixNode {
	return []MatrixNode{}
}

func NewMatrix(rows, cols, length int) *Matrix {
	return &Matrix{
		Nodes:   []MatrixNode{},
		MaxRows: rows,
		MaxCols: cols,
		Length:  length,
	}
}

func (m *Matrix) PushAll(nodes []MatrixNode) {
	m.Nodes = nodes
	sort.Slice(m.Nodes, func(i, j int) bool {
		if m.Nodes[i].row != m.Nodes[j].row {
			return m.Nodes[i].row < m.Nodes[j].row
		} else {
			return m.Nodes[i].col < m.Nodes[j].col
		}
	})
}

func (m *Matrix) Add(n *Matrix) *Matrix {
	p := NewMatrixNodes()
	i := 0
	j := 0
	for i < len(m.Nodes) && j < len(n.Nodes) {
		if m.Nodes[i].row < n.Nodes[j].row {
			p = append(p, m.Nodes[i])
			i++
		} else if m.Nodes[i].row > n.Nodes[j].row {
			p = append(p, n.Nodes[j])
			j++
		} else {
			if m.Nodes[i].col < n.Nodes[j].col {
				p = append(p, m.Nodes[i])
				i++
			} else if m.Nodes[i].col > n.Nodes[j].col {
				p = append(p, n.Nodes[i])
				j++
			} else {
				value := m.Nodes[i].value + n.Nodes[j].value
				if value != 0 {
					p = append(p, MatrixNode{
						row:   m.Nodes[i].row,
						col:   m.Nodes[i].col,
						value: value,
					})
				}
				i++
				j++
			}
		}
	}
	for i < len(m.Nodes) {
		p = append(p, m.Nodes[i])
		i++
	}
	for j < len(n.Nodes) {
		p = append(p, n.Nodes[j])
		j++
	}
	result := NewMatrix(m.MaxRows, m.MaxCols, len(p))
	result.PushAll(p)
	return result
}

func (m *Matrix) Sub(n *Matrix) *Matrix {
	for i := 0; i < len(n.Nodes); i++ {
		n.Nodes[i].value = -n.Nodes[i].value
	}
	return m.Add(n)
}

func main() {
	var rows, cols, length int
	_, _ = fmt.Scanf("%d %d %d\n", &rows, &cols, &length)
	var nodesA = make([]MatrixNode, length)
	for i := 0; i < length; i++ {
		_, _ = fmt.Scanf("(%d,%d,%d)", &nodesA[i].row, &nodesA[i].col, &nodesA[i].value)
	}
	_, _ = fmt.Scanln()
	matrixA := NewMatrix(rows, cols, length)
	matrixA.PushAll(nodesA)
	_, _ = fmt.Scanf("%d %d %d\n", &rows, &cols, &length)
	var nodesB = make([]MatrixNode, length)
	for i := 0; i < length; i++ {
		_, _ = fmt.Scanf("(%d,%d,%d)", &nodesB[i].row, &nodesB[i].col, &nodesB[i].value)
	}
	matrixB := NewMatrix(rows, cols, length)
	matrixB.PushAll(nodesB)
	s1 := matrixA.Add(matrixB)
	fmt.Printf("%d %d %d\n", s1.MaxRows, s1.MaxCols, s1.Length)
	for _, v := range s1.Nodes {
		fmt.Printf("(%d,%d,%d)", v.row, v.col, v.value)
	}
	fmt.Println()
	s1 = matrixA.Sub(matrixB)
	fmt.Printf("%d %d %d\n", s1.MaxRows, s1.MaxCols, s1.Length)
	for _, v := range s1.Nodes {
		fmt.Printf("(%d,%d,%d)", v.row, v.col, v.value)
	}
	fmt.Println()
}
