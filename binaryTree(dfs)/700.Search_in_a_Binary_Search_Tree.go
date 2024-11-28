package binaryTree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    for root != nil && root.Val != val {
        if val > root.Val {
            root = root.Right
        } else {
            root = root.Left
        }
    }
    if root != nil {
        return root
    }
    return nil
}
