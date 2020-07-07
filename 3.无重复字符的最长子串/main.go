package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	l, r := 0, 0
	res := 0
	for r < n {
		c := s[r]
		r++
		m[c]++
		for m[c] > 1 {
			d := s[l]
			l++
			m[d]--
		}
		res = max(res, r+1-l)
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	s := "abcabcbb"
	n := lengthOfLongestSubstring(s)
	fmt.Println(n)
}
