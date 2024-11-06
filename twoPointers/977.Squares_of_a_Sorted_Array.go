// https://leetcode.com/problems/squares-of-a-sorted-array/
package twoPointers
func sortedSquares(nums []int) []int {
    n := len(nums)
    l, r := 0, n-1
    res := make([]int, n)
    for i := n-1; i >= 0; i-- {
        lnum := nums[l]*nums[l]
        rnum := nums[r]*nums[r]

        if lnum > rnum {
            res[i] = lnum
            l++
        } else {
            res[i] = rnum
            r--
        }
    }
    return res
}
