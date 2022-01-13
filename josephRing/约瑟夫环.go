package josephRing

import (
	"fmt"
)

type num struct {
	val, index int
}

func (n num) getVal() int {
	return n.val
}

func (n num) getIndex() int {
	return n.index
}

func main() {
	var n, m int
	_, err := fmt.Scanf("%d %d", &n, &m)
	if err != nil {
		return
	}
	var nums = make([]num, n, n)
	for i := range nums {
		var n int
		_, _ = fmt.Scanf("%d", &n)
		nums[i] = num{
			val:   n,
			index: i + 1,
		}
	}
	mPrint(josephRing(nums, m))
}

func josephRing(nums []num, m int) (ret []int) {
	k := m
	index := 0
	for len(nums) > 0 {
		if k > len(nums) {
			k = k % len(nums)
		}
		index = index + k - 1
		if index >= len(nums) {
			index -= len(nums)
		}
		if index == -1 {
			index = len(nums) - 1
		}
		ret = append(ret, nums[index].getIndex())
		k = nums[index].getVal()
		if index != len(nums)-1 {
			nums = append(nums[:index], nums[index+1:]...)
		} else {
			nums = nums[:index]
		}
		if index == len(nums) {
			index = 0
		}
	}
	return
}

func mPrint(nums []int) {
	for _, v := range nums {
		fmt.Printf("%d ", v)
	}
}

/*
package main

import "fmt"
func panduan(num []int)int64{
	var ans int64
	for _,v:=range num{
		if v!=0{
			ans++
		}
	}
	return ans
}
func main() {
	var (
		num []int
		n int
		m int
		t int
	)
	fmt.Scan(&n,&m)
	for i:=0;i<n;i++{
		fmt.Scan(&t)
		num=append(num,t)
	}
	i:=0
	ans:=1
	for;panduan(num)>=1;{
		if i>len(num)-1{
			i=0
		}
		if num[i]==0{
			i++
			continue
		}
		if ans==m{
			m=num[i]
			fmt.Printf("%d ",i+1)
			num[i]=0
			ans=0
		}
		i++
		ans++
	}
}
*/
