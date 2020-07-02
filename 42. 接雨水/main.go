func trap(height []int) int {
    left, right := 0, len(height)-1
    l_max, r_max := 0, 0
    res := 0

    for left < right {
        if height[left] < height[right] {
            if height[left] > l_max {
                l_max = height[left]
            } else {
                res += l_max - height[left]
            }
            left++
        } else {
            if height[right] > r_max {
                r_max = height[right]
            } else {
                res += r_max - height[right]
            }
            right--
        }
    }
    
    return res
}