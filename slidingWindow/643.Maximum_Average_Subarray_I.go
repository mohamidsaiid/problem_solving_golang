// https://leetcode.com/problems/maximum-average-subarray-i/
package slidingWindow
func findMaxAverage(nums []int, k int) float64 {
    var sum float64
    for i := 0; i < k; i++ {
        sum += float64(nums[i])
    }
    maxx := sum/float64(k)
    l, r := 0, k
    for r < len(nums) {
        sum -= float64(nums[l])
        sum += float64(nums[r])
        l++
        r++
        maxx = max(maxx, sum/float64(k))
    }
    return maxx
}
