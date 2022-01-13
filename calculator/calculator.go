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
	val   int
	isNum bool
}

//计算中缀
func main() {
	var s []Num
	for {
		var c1, c2 int
		_, err := fmt.Scanf("%d%c", &c1, &c2)
		if err != nil {
			fmt.Println(err)
			return
		}
		s = append(s, Num{
			val:   c1,
			isNum: true,
		})
		if c2 == '#' {
			break
		}
		s = append(s, Num{
			val:   c2,
			isNum: false,
		})
	}
	fmt.Println(CalculateTheSuffix(InfixToSuffix(s)))
}

func ToNum(str string) (num []Num) {
	for i := 0; i < len(str); i++ {
		num = append(num, Num{
			val:   int(str[i]),
			isNum: IsNum(str[i]),
		})
	}
	return
}

func IsNum(c byte) bool {
	return c >= '0' && c <= '9'
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
					if ch.Val() != '(' && Priority[s[i].Val()] <= Priority[ch.Val()] {
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

func (n Num) IsNum() bool {
	return n.isNum
}

func (n Num) Val() int {
	return n.val
}

func CalculateTheSuffix(s []Num) int {
	numStack := CreatStack()
	for i := 0; i < len(s); i++ {
		if s[i].IsNum() {
			numStack.Push(s[i])
		} else {
			right := numStack.Pop()
			left := numStack.Pop()
			numStack.Push(Calculate(left.Val(), right.Val(), s[i].Val()))
		}
	}
	return numStack.Pop().Val()
}

func Calculate(left, right, operate int) Num {
	num := Num{
		isNum: true,
	}
	switch operate {
	case '+':
		num.val = left + right
	case '-':
		num.val = left - right
	case '*':
		num.val = left * right
	case '/':
		num.val = left / right
	}
	return num
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
