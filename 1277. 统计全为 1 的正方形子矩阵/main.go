func countSquares(matrix [][]int) int {
    m := len(matrix)
    n := len(matrix[0])
    dp := make([][]int, m)
    res := 0

    for i:=0; i<m; i++ {
        dp[i] = make([]int, n)
        for j:=0; j<n; j++ {
            if i==0 || j==0 {
                dp[i][j] = matrix[i][j]
            } else if matrix[i][j] == 1 {
                dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
            } else {
                dp[i][j] = 0
            }
            res += dp[i][j]
        }
    }

    return res
}

func min(a, b, c int) int {
    min := a
    if min > b {
        min = b
    }
    if min > c {
        min = c
    }
    return min
}