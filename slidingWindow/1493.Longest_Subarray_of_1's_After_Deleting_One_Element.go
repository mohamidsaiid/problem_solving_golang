// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/
package slidingWindow
func longestSubarray(nums []int) int {
    k := 1
    maxLen := 0
    l := 0
    zeros := 0
    for r := 0; r < len(nums); r++ {
        if nums[r] == 0 { zeros++; }
        for zeros > k {
            if nums[l] == 0 { zeros--; }
            l++
        }
        maxLen = max(maxLen, r - l )
    }
    return maxLen
}
