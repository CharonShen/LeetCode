func levelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)
    if root == nil {
        return res
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)

    for len(queue) > 0 {
        list := make([]int, 0)
        l := len(queue)
        for i:=0; i<l; i++ {
            temp := queue[0]
            queue = queue[1:]
            list = append(list, temp.Val)
            if temp.Left != nil {
                queue = append(queue, temp.Left)
            }
            if temp.Right != nil {
                queue = append(queue, temp.Right)
            }
        }
        res = append(res, list)
    }
    return res
}