func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }
    if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else {
        if root.Left == nil {
            return root.Right
        } else if root.Right == nil {
            return root.Left
        } else {
            temp := root.Right
            for temp.Left != nil {
                temp = temp.Left
            }
            temp.Left = root.Left
            return root.Right
        }
    }
    return root
}