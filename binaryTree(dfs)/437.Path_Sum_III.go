package binaryTree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    cnt := 0

    var dfsSum func(t *TreeNode, sum int)
    dfsSum = func(t *TreeNode, sum int) {
        if t == nil {
            return
        }
        if sum + t.Val == targetSum {
            cnt++   
        }
        dfsSum(t.Left, sum+t.Val) 
        dfsSum(t.Right, sum+t.Val)
    }

    var dfs func(t *TreeNode) 
    dfs = func(t *TreeNode) {
        if t == nil {
            return 
        }
        dfsSum(t, 0)
        dfs(t.Left)
        dfs(t.Right)
    }
    dfs(root)
    return cnt
}
