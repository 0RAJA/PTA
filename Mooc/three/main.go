package main

import "fmt"

func replaceStr(r string, ch1, ch2 rune) (str string) {
	for _, v := range r {
		if v == ch1 {
			str += string(ch2)
		} else {
			str += string(v)
		}
	}
	return
}

func deleteStr(r string, ch rune) (str string) {
	for _, v := range r {
		if v != ch {
			str += string(v)
		}
	}
	return
}

func startCompare(r1 string, index int, r2 string) int {
	offset := kmp(r1[index-1:], r2)
	if offset == -1 {
		return -1
	}
	return index + offset - 1
}

func kmp(s1, s2 string) int {
	next := getNext(s2)
	for i, j := 0, 0; i < len(s1); {
		if j == -1 || s1[i] == s2[j] {
			i++
			j++
		} else {
			j = next[j]
		}
		if j == len(s2) {
			return i - j
		}
	}
	return -1
}

func getNext(s string) []int {
	next := make([]int, len(s))
	next[0] = -1
	for i, j := 0, -1; i < len(s)-1; {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			if s[i] == s[j] {
				next[i] = next[j]
			} else {
				next[i] = j
			}
		} else {
			j = next[j]
		}
	}
	return next
}

func StrReplace(S, T, V string) (ret string) {
	index := 0
	for {
		x := startCompare(S, index+1, T)
		if x == -1 {
			return
		}
		ret += S[index:x] + V
		index = x + len(T)
	}
}

func main() {
	fmt.Println(replaceStr("aaaaabbbccc", 'a', 'z'))   //zzzzzbbbccc
	fmt.Println(deleteStr("aaaaabbbccc", 'a'))       		   //bbbccc
	fmt.Println(startCompare("abcdef", 1, "ef"))     // 4  从下标为4开始匹配
	fmt.Println(startCompare("abcdef", 2, "bcd"))    // 1  从下标为2开始匹配成功
	fmt.Println(startCompare("abcdef", 3, "bcd"))    // -1 匹配不成功
	fmt.Println(StrReplace("abcaaaabc", "abc", "cba"))    // cbaaaacba 将abcaaaabc中的abc替换为cba
}
