package main

import "fmt"

type num struct {
	val, index int
	next       *num
}

func main() {
	var n, m int
	_, err := fmt.Scanf("%d %d", &n, &m)
	if err != nil {
		return
	}
	var nums = make([]int, n)
	for i := range nums {
		//_, _ = fmt.Scan(&nums[i])
		nums[i] = m
	}
	tail := creatLink(nums)
	mPrint(run(m, tail, n))
}

func creatLink(nums []int) *num {
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
		//length--
		n = t.getVal()
		//if n == 0 {
		//	n = length
		//}
		ret = append(ret, t.getIndex())
	}
	return append(ret, node.getIndex())
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
