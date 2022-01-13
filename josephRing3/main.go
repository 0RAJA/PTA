package main

import (
	"fmt"
)

type num struct {
	val, index int
	next       *num
}

func main() {
	var n, k int
	_, err := fmt.Scanf("%d %d", &n, &k)
	if err != nil {
		return
	}
	model := make([]int, k)
	for i := range model {
		_, _ = fmt.Scan(&model[i])
	}
	for i := 1; i <= n; i++ {
		tail := creatLink(n)
		x := run(i, tail, n)
		//fmt.Println(x)
		if IsEqual(k, x, model) {
			fmt.Print(i)
		}
	}
}

func creatLink(n int) *num {
	nums := make([]int, n)
	var head, tail *num
	for i := range nums {
		p := &num{
			val:   nums[i],
			index: i + 1,
		}
		if head == nil {
			head = p
			head.next = head
			tail = head
		} else {
			p.next = tail.next
			tail.next = p
			tail = p
		}
	}
	return tail
}

func run(n int, node *num, length int) (ret []int) {
	ret = make([]int, 0, length)
	for node.next != node {
		for i := 1; i < n; i++ {
			node = node.next
		}
		t := node.next
		node.next = node.next.next
		length--
		if n == 0 {
			n = length
		}
		ret = append(ret, t.getIndex())
	}
	return append(ret, node.getIndex())
}
func IsEqual(k int, nums, model []int) bool {
	for i := 0; i < k; i++ {
		if nums[len(nums)-i-1] != model[len(model)-i-1] {
			return false
		}
	}
	return true
}
func (n *num) getVal() int {
	return n.val
}

func (n *num) getIndex() int {
	return n.index
}

func mPrint(nums []int) {
	for _, v := range nums {
		fmt.Printf("%d ", v)
	}
}
