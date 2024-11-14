// https://leetcode.com/problems/find-pivot-index/
package prefixSum
func pivotIndex(nums []int) int {
    size := len(nums)
    pre, suf := make([]int, size), make([]int, size)

    for i := 1; i < size; i++ {
        pre[i] = pre[i-1] + nums[i-1]
    }
    for i := size-2; i >= 0; i-- {
        suf[i] = suf[i+1] + nums[i+1]
    }
    for i := 0; i < size; i++ {
        if suf[i] == pre[i] {
            return i
        }
    }
    return -1
}
