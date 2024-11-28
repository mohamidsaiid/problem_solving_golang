package binaryTree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    leafs1, leafs2 := make([]int, 0, 200), make([]int, 0, 200)
    var dfs func(t *TreeNode, leafs *[]int)
    dfs = func(t *TreeNode, leafs *[]int) {
        if t == nil {
            return
        }
        if t.Right == nil && t.Left == nil {
            *leafs = append(*leafs, t.Val)
            return 
        }
        dfs(t.Left, leafs)
        dfs(t.Right, leafs)
    }

    dfs(root1, &leafs1)
    dfs(root2, &leafs2)
    if len(leafs1) != len(leafs2) {
        return false
    }
    size := len(leafs1)
    for i := 0; i < size; i++ {
        if leafs1[i] != leafs2[i] {
            return false
        }
    }
    return true
}
