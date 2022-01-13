package main

import "fmt"

var stack []byte

func main() {
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		fmt.Println(err)
		return
	}
	if IsPalindrome(s) {
		fmt.Println("回文")
	}else {
		fmt.Println("不是回文")
	}
}

func IsPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	mid := len(s) / 2
	for i := 0; i < mid; i++ {
		Push(s[i])
	}
	if len(s)%2 != 0 {
		mid++
	}
	for len(stack) > 0 {
		ch := Pop()
		if ch == s[mid] {
			mid++
		} else {
			return false
		}
	}
	return true
}

func Push(ch byte) {
	stack = append(stack, ch)
}

func Pop() (ch byte) {
	ch = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return
}
