// https://leetcode.com/problems/max-number-of-k-sum-pairs/

package twoPointers
import "slices"
func maxOperations(nums []int, k int) int {
    l, r := 0, len(nums) - 1
    res := 0
    slices.Sort(nums)
    for l < r {
        if nums[l] + nums[r] < k {
            l++
        } else if nums[l] + nums[r] > k {
            r--
        } else {
            res++
            l++
            r--
        }
    }
    return res
}
