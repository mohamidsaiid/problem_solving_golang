package binaryTree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    
    goodNodes := 0
    var dfs func(t *TreeNode, maxNode int) 
    dfs = func(t *TreeNode, maxNode int) {
        if t == nil {
            return
        }
        if t.Val >= maxNode {
            goodNodes++
            maxNode = t.Val
        }
        dfs(t.Left, maxNode)
        dfs(t.Right, maxNode)
    }
    dfs(root, root.Val)
    return goodNodes
}
