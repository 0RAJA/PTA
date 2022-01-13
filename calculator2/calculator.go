package main

import (
	"fmt"
)

var Priority = map[int]int{
	'+': 0,
	'-': 0,
	'*': 1,
	'/': 1,
}

type Stack struct {
	stack []Num
}

type Num struct {
	val   byte
	isNum bool
}

//中缀转后缀
func main() {
	var s []Num
	for {
		var c byte
		_, _ = fmt.Scanf("%c", &c)
		if c == '#' {
			break
		}
		s = append(s, Num{
			val:   c,
			isNum: IsNum(c),
		})
	}
	Print(InfixToSuffix(s))
}

// InfixToSuffix 中缀转后缀
func InfixToSuffix(s []Num) (ret []Num) {
	charStack := CreatStack() //字符栈
	for i := 0; i < len(s); i++ {
		if s[i].IsNum() {
			ret = append(ret, s[i])
		} else {
			switch s[i].Val() {
			case '(':
				charStack.Push(s[i])
			case ')':
				for !charStack.IsEmpty() {
					ch := charStack.Pop()
					if ch.Val() == '(' {
						break
					}
					ret = append(ret, ch)
				}
			default:
				for !charStack.IsEmpty() {
					ch := charStack.Pop()
					if ch.Val() != '(' && Priority[int(s[i].Val())] <= Priority[int(ch.Val())] {
						ret = append(ret, ch)
					} else {
						charStack.Push(ch)
						break
					}
				}
				charStack.Push(s[i])
			}
		}
	}
	for !charStack.IsEmpty() {
		ret = append(ret, charStack.Pop())
	}
	return
}

func IsNum(n byte) bool {
	return n >= 'a' && n <= 'z'
}

func (n Num) IsNum() bool {
	return n.isNum
}

func (n Num) Val() byte {
	return n.val
}

func Print(nums []Num) {
	for i := range nums {
		fmt.Print(string(nums[i].Val()))
	}
}

func CreatStack() *Stack {
	return &Stack{stack: make([]Num, 0)}
}

func (s *Stack) Push(ch Num) {
	s.stack = append(s.stack, ch)
}

func (s *Stack) Pop() (ch Num) {
	ch = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}
