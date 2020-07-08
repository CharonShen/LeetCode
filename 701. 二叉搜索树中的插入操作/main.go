func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        node := new(TreeNode)
        node.Val = val
        return node
    }
    if root.Val > val {
        root.Left = insertIntoBST(root.Left, val)
    } else {
        root.Right = insertIntoBST(root.Right, val)
    }
    return root
}