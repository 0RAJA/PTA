package two

import (
	"fmt"
	"testing"
)

func AddNode(l, r int) *Node {
	var head *Node
	for i := l; i <= r; i++ {
		p := &Node{
			val: i,
		}
		if head == nil {
			head = p
		} else {
			p.next = head.next
			head.next = p
		}
	}
	return head
}

func PrintNode(head *Node) {
	for p := head; p != nil; p = p.next {
		fmt.Print(p.val, " ")
	}
	fmt.Println()
}

func TestReverse2(t *testing.T) {
	head := AddNode(1, 2)
	PrintNode(head)
	fmt.Println()
	Reverse2(head)
	PrintNode(head)
}

func TestDelAndInsert(t *testing.T) {
	la := AddNode(1, 10)
	lb := AddNode(11, 20)
	PrintNode(la)
	PrintNode(lb)
	DelAndInsert(la, lb, 2, 2, 5)
	PrintNode(la)
	PrintNode(lb)
}
