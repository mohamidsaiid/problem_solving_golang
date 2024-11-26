package linkedList
type ListNode struct {
     Val int
     Next *ListNode
}
func deleteMiddle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }

    prev, slow, fast := head, head, head
    for slow != nil && fast != nil && fast.Next != nil {
        prev, slow, fast = slow, slow.Next, fast.Next.Next
    }
    prev.Next = slow.Next
    return head
}
