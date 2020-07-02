func inorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    stack := make([]*TreeNode, 0)
    for len(stack) > 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        temp := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, temp.Val)
        root = temp.Right
    }
    return res
}