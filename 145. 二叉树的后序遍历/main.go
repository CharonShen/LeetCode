package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastNode *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		temp := stack[len(stack)-1]
		if temp.Right == nil || temp.Right == lastNode {
			stack = stack[:len(stack)-1]
			res = append(res, temp.Val)
			lastNode = temp
		} else {
			root = temp.Right
		}
	}
	return res
}
func main() {
	root3 := &TreeNode{
		Val: 3,
	}
	root2 := &TreeNode{
		Val:  2,
		Left: root3,
	}
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: root2,
	}
	res := postorderTraversal(root)
	fmt.Println(res)
}
