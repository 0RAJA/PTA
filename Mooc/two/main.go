package two

/*
设有一线性表e=(e1 , e2 , … , en-1 , en)，其逆线性表定义为e'=( en , en-1 , … , e2 , e1)，
请设计一个算法，将线性表逆置，要求逆线性表仍占用原线性表的空间，并且用顺序表和单链表两种方法来表示，写出不同的处理函数。
*/

type Node struct {
	val  int
	next *Node
}

func Reverse1(nums []int) {
	for i := 0; i < len(nums); i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}

// Reverse2 带头结点
func Reverse2(head *Node) {
	p := head.next
	q := head
	for p != nil {
		t := p.next
		p.next = q
		q = p
		p = t
	}
	if head.next != nil {
		head.next.next = nil
	}
	head.next = q
}

/*
‌设指针la和lb分别指向两个无头结点单链表中的首元结点，
试设计从表la中删除自第i个元素起共len个元素，并将它们插入到表lb的第j个元素之后的算法。
*/

func DelAndInsert(la, lb *Node, i, j, length int) (newA *Node, newB *Node) {
	if la == nil {
		return nil, lb
	}
	pA := &Node{next: la}
	for m := 0; pA.next != nil && m < i-1; m++ {
		pA = pA.next
	}
	pB := pA.next
	for m := 1; pB.next != nil && m < length; m++ {
		pB = pB.next
	}
	t := pA.next
	if i == 1 {
		newA = pB.next
	} else {
		pA.next = pB.next
		newA = pA
	}
	pA = t
	pC := &Node{next: lb}
	for m := 0; pC.next != nil && m < j-1; m++ {
		pC = pC.next
	}
	pB.next = pC.next
	pC.next = pA
	newB = lb
	return
}
