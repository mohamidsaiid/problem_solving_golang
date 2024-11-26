package linkedList
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    nhead := new(ListNode)
    nhead.Val = head.Val
    for head.Next != nil {
        tmpHead := new(ListNode)
        tmpHead.Val = head.Next.Val
        tmpHead.Next = nhead
        nhead = tmpHead
        head = head.Next
    }
    return nhead
}
