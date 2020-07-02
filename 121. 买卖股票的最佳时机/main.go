func maxProfit(prices []int) int {
    ans, n := 0, len(prices)
    if n == 0 {
        return 0
    }
    
    minVal := prices[0]

    for i:=0; i<n; i++ {
        if prices[i] < minVal {
            minVal = prices[i]
        } else if ans < prices[i] - minVal {
            ans = prices[i] - minVal
        }
    }
    
    return ans
}