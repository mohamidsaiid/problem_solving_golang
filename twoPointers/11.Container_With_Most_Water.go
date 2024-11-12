// https://leetcode.com/problems/container-with-most-water/
package twoPointers
func maxArea(height []int) int {
    l, r := 0, len(height) - 1
    res := 0
    for l < r {
        width := r - l
        if height[l] <= height[r] {
            res = max(res, height[l]*width)
            l++
        } else {
            res = max(res, height[r]*width)
            r--
        }
    }
    return res
}
