package linkedList
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    oddHead, evenHead := head, head.Next
    odd, even := oddHead, evenHead

    for even != nil && even.Next != nil {
        odd.Next = even.Next
        even.Next = even.Next.Next
        odd = odd.Next
        even = even.Next
    }

    odd.Next = evenHead
    return oddHead
}
