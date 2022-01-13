package main

import "fmt"

func cycleRight(nums []int, k int) (newNums []int) {
	k %= len(nums)
	getNext := func(now, k, length int) (ret int) {
		ret = ((now - k) + length) % length
		return
	}
	temp := nums[0]
	nx := 0
	for {
		p := getNext(nx, k, len(nums))
		if p == 0 {
			break
		}
		nums[nx] = nums[p]
		nx = p
	}
	nums[nx] = temp
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(cycleRight(nums, 3)) // [4 2 3 1 5 6]
}
