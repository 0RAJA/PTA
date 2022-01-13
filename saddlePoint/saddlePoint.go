package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	row, col, value int
}

func main() {
	var row, col int
	_, _ = fmt.Scanf("%d %d\n", &row, &col)
	nums := make([][]int, row+1)
	for i := 0; i < len(nums); i++ {
		nums[i] = make([]int, col+1)
	}
	buf := bufio.NewReader(os.Stdin)
	for i := 1; i <= row; i++ {
		s, _ := buf.ReadString('\n')
		s = strings.Replace(s, "\n", "", -1)
		strs := strings.Split(s, " ")
		for j := 1; j <= col; j++ { //初始化,将nums[i][0] = 第i行中最小,nums[0][j] = 第j列中最大
			nums[i][j], _ = strconv.Atoi(strs[j-1])
			if i == 1 {
				nums[0][j] = nums[i][j]
			} else if nums[0][j] < nums[i][j] {
				nums[0][j] = nums[i][j]
			}
			if j == 1 {
				nums[i][0] = nums[i][j]
			} else if nums[i][0] > nums[i][j] {
				nums[i][0] = nums[i][j]
			}
		}
	}
	var ret []Node
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			if nums[i][j] == nums[i][0] && nums[i][j] == nums[0][j] { //寻找鞍点
				ret = append(ret, Node{
					row:   i,
					col:   j,
					value: nums[i][j],
				})
			}
		}
	}
	for _, v := range ret {
		fmt.Printf("(%d,%d,%d)", v.row, v.col, v.value)
	}
	if len(ret) == 0 {
		fmt.Printf("NONE")
	}
}
