package linkedList
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    copyList := func(p *ListNode) *ListNode {
        nHead := new(ListNode)
        nHead.Val = p.Val
        tmp := nHead
        for p.Next != nil {
            tmp.Next = new(ListNode)
            tmp.Next.Val = p.Next.Val
            tmp = tmp.Next
            p = p.Next
        }
        return nHead
    }
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
    
    revHead := rev(copyList(head))
    maxSum := 0
    for head != nil && revHead != nil {
        maxSum = max(revHead.Val + head.Val, maxSum)
        revHead = revHead.Next
        head = head.Next
    }
    return maxSum
}
