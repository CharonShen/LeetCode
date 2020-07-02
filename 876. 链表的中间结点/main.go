func middleNode(head *ListNode) *ListNode {
    fast, slow := head, head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    // if fast.Next != nil {
    //     slow = slow.Next
    // }
    return slow
}