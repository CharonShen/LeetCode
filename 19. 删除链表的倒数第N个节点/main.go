func removeNthFromEnd(head *ListNode, n int) *ListNode {
    temp := head
    s := 1
    for {
        if temp.Next == nil {
            break
        }
        temp = temp.Next
        s++
    }
    temp = head
    for i:=0; i<s-n-1; i++ {
        temp = temp.Next
    }
    if s == 1 || n == s {//删除头节点
        head = head.Next
    } else if n == 1 {//删除尾节点
        temp.Next = nil
    } else {
        temp.Next = temp.Next.Next
    }
    return head
}