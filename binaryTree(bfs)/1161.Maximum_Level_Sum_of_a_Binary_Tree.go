package binaryTree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    maxLevelSum := root.Val

    currentLevel, maxLevel := 0, 1
    
    queue := make([]*TreeNode, 1, 100)
    queue[0] = root
    
    for len(queue) != 0 {
        currentLevel++
        lvlSize := len(queue)
        sumOfLevel := 0
        for lvlSize > 0 {
            lvlSize--
            node := queue[0]
            queue = queue[1:]
            sumOfLevel += node.Val
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
        }
        if maxLevelSum < sumOfLevel {
            maxLevelSum = sumOfLevel
            maxLevel = currentLevel
        }
    }
    return maxLevel
}
