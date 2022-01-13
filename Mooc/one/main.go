package one

/*
设计求解下列问题的算法，并分析其最坏情况的时间复杂度及其量级。
（1）在数组A[1..n]中查找值为K的元素，若找到则输出其位置i(1<=i<=n)，否则输出0作为标志。
（2）找出数组A[1..n]中元素的最大值和次最大值（本小题以数组元素的比较为标准操作）。
*/

func Find1(nums []int, K int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] == K {
			return i
		}
	}
	return 0
}

func Find2(nums []int) (min int, max int) {
	min = nums[1]
	max = nums[1]
	for i := 2; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return
}
