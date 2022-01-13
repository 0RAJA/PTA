package main

import (
	"fmt"
	"time"
)

const (
	ROW = 8
	COL = 8
)

type Stack struct {
	x, y      int
	direction int
}

var (
	next       = [8][2]int{{-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}}
	stack      [ROW*COL + 1]Stack
	checkBoard [ROW][COL]int
	index      = 1 //初始栈下标
)

// Checkerboard 马踏棋盘
func Checkerboard(x, y int) {
	Push(Stack{
		x:         x,
		y:         y,
		direction: 0,
	})
Loop:
	for index <= ROW*COL {
		p := Front()
		for p.direction < len(next) {
			nx, ny := p.x+next[p.direction][0], p.y+next[p.direction][1]
			p.direction++
			if IsLegal(nx, ny) {
				Push(Stack{
					x:         nx,
					y:         ny,
					direction: 0,
				})
				goto Loop
			}
		}
		Pop()
	}
}

func main() {
	var sx, sy int
	_, _ = fmt.Scanf("%d %d", &sx, &sy)
	t1 := time.Now()
	if sx <= 0 || sx > ROW || sy <= 0 || sy > COL {
		fmt.Println("Border Error")
		return
	}
	Checkerboard(sx-1, sy-1)
	fmt.Println(time.Since(t1))
	Print()

}

func IsLegal(x, y int) bool {
	if x >= 0 && y >= 0 && x < ROW && y < COL && checkBoard[x][y] == 0 {
		return true
	}
	return false
}

func Print() {
	for i := 0; i < len(checkBoard); i++ {
		for j := 0; j < len(checkBoard); j++ {
			fmt.Printf("%2d ", checkBoard[i][j])
		}
		fmt.Println()
	}
}

func Push(p Stack) {
	stack[index] = p
	checkBoard[p.x][p.y] = index
	index++
}

func Front() *Stack {
	return &(stack[index-1])
}

func Pop() {
	p := Front()
	index--
	checkBoard[p.x][p.y] = 0
}
