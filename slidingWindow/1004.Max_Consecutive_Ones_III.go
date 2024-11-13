// https://leetcode.com/problems/max-consecutive-ones-iii/

package slidingWindow
func longestOnes(nums []int, k int) int {
    maxLen := 0
    l := 0
    zeros := 0
    for r := 0; r < len(nums); r++ {
        if nums[r] == 0 { zeros++; }
        for zeros > k {
            if nums[l] == 0 { zeros--; }
            l++
        }
        maxLen = max(maxLen, r - l + 1)
    }
    return maxLen
}
