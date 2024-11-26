package linkedList
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseListRecursive(head *ListNode) *ListNode {
    var rev func(p *ListNode) *ListNode
    rev = func(p *ListNode) *ListNode {
        if p == nil || p.Next == nil {
            return p
        }
        n := rev(p.Next)
        p.Next.Next = p
        p.Next = nil
        return n
    }

    nhead := rev(head)
    return nhead
}
