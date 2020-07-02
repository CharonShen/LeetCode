package main

import "fmt"

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for k := 0; k < m; k++ {
		dp[k] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i != 0 || j != 0 {
				if i == 0 {
					dp[i][j] = dp[i][j-1]
				} else if j == 0 {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i-1][j] + dp[i][j-1]
				}
			}

		}
	}
	res := dp[m-1][n-1]
	return res
}
func main() {
	res := uniquePaths(3, 2)
	fmt.Println(res)
}
