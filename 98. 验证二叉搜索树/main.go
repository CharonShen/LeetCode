// func helper(root *TreeNode, low, high int) bool {
//     if root == nil {
//         return true
//     }
//     if root.Val <= low || root.Val >= high {
//         return false
//     }
//     return helper(root.Left, low, root.Val) && helper(root.Right, root.Val, high)
// }
func isValidBST(root *TreeNode) bool {
    // return helper(root, math.MinInt64, math.MaxInt64)
    lastVal := math.MinInt64
    stack := make([]*TreeNode, 0)
    
    for len(stack) != 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        temp := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if temp.Val <= lastVal {
            return false
        }
        lastVal = temp.Val
        root = temp.Right
    }
    return true
}