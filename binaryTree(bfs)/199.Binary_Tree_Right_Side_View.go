package binaryTree
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func rightSideView(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    res := make([]int, 0, 100)
    
    queue := make([]*TreeNode, 1, 100)
    queue[0] = root
    
    for len(queue) != 0 {
        res = append(res, queue[0].Val)
        tmpQ := queue[:]
        for len(tmpQ) != 0 {
            node := tmpQ[0]
            tmpQ = tmpQ[1:]
            queue = queue[1:]
            if node != nil {
                if node.Right != nil {
                    queue = append(queue, node.Right)
                }
                if node.Left != nil {
                    queue = append(queue, node.Left)
                }
            }
        }
    }
    return res
}
