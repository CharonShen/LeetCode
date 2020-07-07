package main

import "fmt"

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	start, maxLen := 0, 0
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i > maxLen {
				start = i
				maxLen = j - i
			}
		}
	}
	return s[start : start+maxLen+1]
}

func main() {
	str := "cbbd"
	res := longestPalindrome(str)
	fmt.Println(res)
}
