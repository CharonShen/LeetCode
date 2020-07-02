func minPathSum(grid [][]int) int {
    m := len(grid)
    if m == 0 {
        return 0
    }
    n := len(grid[0])
    dp := make([][]int, m)
    res := 0

    for i:=0; i<m; i++ {
        dp[i] = make([]int, n)
        if i==0 {
            dp[0][0] = grid[0][0]
        }
        for j:=0; j<n; j++ {
            if i!=0 || j!=0 {
                if i==0 {
                    dp[i][j] = dp[i][j-1] + grid[i][j]
                } else if j==0 {
                    dp[i][j] = dp[i-1][j] + grid[i][j]
                } else {
                    dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
                }
            }
        }
    }
    res = dp[m-1][n-1]
    return res
}
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}