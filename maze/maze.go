package main

import (
	"fmt"
)

type Stack struct {
	x, y      int //当前状态的行列
	direction int //当前方向
}

const (
	SX = 1
	SY = 1
)

var (
	Next     = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	row, col int     //迷宫行列
	stack    []Stack //栈
	maze     [][]int //记录地图信息
)

func Maze() {
	Push(Stack{
		x:         SX,
		y:         SY,
		direction: 0,
	})
Loop:
	for len(stack) > 0 {
		p := Front() //获取栈顶元素
		for p.direction < len(Next) {
			nx, ny := p.x+Next[p.direction][0], p.y+Next[p.direction][1]
			p.direction++
			if IsLegal(nx, ny) { //判断nx,ny是否超过边界以及是否已经走过
				Push(Stack{
					x:         nx,
					y:         ny,
					direction: 0,
				})
				if nx == row && ny == col { //到达目标点
					Print()
					return
				}
				goto Loop
			}
		}
		Pop()
	}
}

func main() {
	_, err := fmt.Scanf("%d %d", &row, &col)
	if err != nil {
		fmt.Println(err)
		return
	}
	maze = make([][]int, row+1)
	for i := range maze {
		maze[i] = make([]int, col+1)
	}
	for i := 1; i <= row; i++ {
		var s string
		_, err := fmt.Scanln(&s)
		if err != nil {
			fmt.Println(err)
			return
		}
		for j := 1; j <= col; j++ {
			maze[i][j] = int(s[j-1] - '0')
		}
	}
	Maze()
}

func IsLegal(x, y int) bool {
	return x > 0 && y > 0 && x <= row && y <= col && maze[x][y] == 0
}

func Push(s Stack) {
	maze[s.x][s.y] = -1
	stack = append(stack, s)
}

func Pop() {
	p := stack[len(stack)-1]
	maze[p.x][p.y] = 0
	stack = stack[:len(stack)-1]
}

func Front() *Stack {
	return &stack[len(stack)-1]
}

func Print() {
	for i := 0; i < len(stack); i++ {
		if i == len(stack)-1 {
			stack[i].direction = 0
		}
		fmt.Printf("(%d,%d,%d)", stack[i].x, stack[i].y, stack[i].direction)
	}
}
