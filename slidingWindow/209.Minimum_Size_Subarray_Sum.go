// https://leetcode.com/problems/minimum-size-subarray-sum/
package slidingWindow
func minSubArrayLen(target int, nums []int) int {
    n := len(nums)

    l := 0
    sum := 0
    minWinSize := math.MaxInt32
    flag := false

    for r := 0; r < n; r++ {
        sum += nums[r]
        for sum >= target {
            minWinSize = min(minWinSize, r - l+1)
            flag = true
            sum -= nums[l]
            l++
        }         
    }
    if !flag {
        return 0
    }
    return minWinSize
}
