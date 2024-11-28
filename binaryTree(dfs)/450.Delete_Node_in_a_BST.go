package binaryTree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return root
    } else if root.Val  < key {
        root.Right = deleteNode(root.Right, key)
    } else if root.Val > key {
        root.Left = deleteNode(root.Left, key)
    } else {
        if root.Left == nil && root.Right == nil {
            return nil
        } else if root.Left == nil {
            return root.Right
        } else if root.Right == nil {
            return root.Left
        } else {
            root.Val = findMin(root.Right)
            root.Right = deleteNode(root.Right, root.Val)
        }
    }
    return root
}
func findMin(n *TreeNode) int {
    minV := n.Val
    for n != nil {
        if n.Val < minV {
            minV = n.Val
        }
        n = n.Left
    }
    return minV
}
