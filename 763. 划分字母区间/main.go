package main

import "fmt"

func partitionLabels(S string) []int {
	ans := []int{}
	l := len(S)
	if l <= 1 {
		ans = append(ans, l)
		return ans
	}

	m := make(map[byte]int)
	for i := 0; i < l; i++ {
		m[S[i]] = i
	}
	i := 0
	j := m[S[i]]
	for i < l {
		flag := true
		for k := i + 1; k < j; k++ {
			if m[S[k]] > j {
				flag = false
				j = max(j, m[S[k]])
			}
		}
		if flag {
			ans = append(ans, j-i+1)
			i = j + 1
			if i < l {
				j = m[S[i]]
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	s := "ababcbacadefegdehijhklij"
	res := partitionLabels(s)
	fmt.Println(res)
}
