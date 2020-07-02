func lengthOfLIS(nums []int) int {
    n := len(nums)
    if n == 0 || n == 1 {
        return n
    }
    dp := make([]int, n)
    dp[0] = 1
    res := 1
    for i:=1; i<n; i++ {
       for j:=0; j<i; j++ {
           cur := 1
           if nums[i] > nums[j] {
               cur = dp[j]+1
           }
           dp[i] = max(dp[i], cur)
       }
       res = max(res, dp[i])
    }
    return res
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}