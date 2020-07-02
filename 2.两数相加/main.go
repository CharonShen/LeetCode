/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 = new(ListNode)
    current := l3
    flag := 0
    for l1 != nil || l2 != nil || flag > 0 {
        x, y := 0, 0
        if l1 != nil {
            x = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            y = l2.Val
            l2 = l2.Next
        }
        sum := x + y + flag
        current.Val = sum%10
        flag = sum/10
        if l1 != nil || l2 != nil || flag > 0{
            current.Next = new(ListNode)
            current = current.Next
        }
    }
    return l3
}