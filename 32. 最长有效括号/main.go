func longestValidParentheses(s string) int {
    stack := []int{}
    max := 0
    stack = append(stack, -1)
    for i, v := range s {
        if v == '(' {
            stack = append(stack, i)
        } else {
            stack = stack[:len(stack)-1]
            if len(stack) == 0 {
                stack = append(stack, i)
            }
            max = Max(max, i-stack[len(stack)-1])
        }
    }
    return max
}
func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}