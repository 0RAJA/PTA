package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println(err)
		return
	}
	triangle := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		triangle[i] = make([]int, i+2)
		triangle[i][1] = 1
		for j := 2; j <= i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", triangle[i][j])
		}
		fmt.Println()
	}
}
