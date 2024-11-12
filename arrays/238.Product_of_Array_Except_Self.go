// https://leetcode.com/problems/product-of-array-except-self/
package arrays
func productExceptSelf(nums []int) []int {
    size := len(nums)
    pre, suf := make([]int, size), make([]int, size)
    res := make([]int, size)
    pre[0] = 1
    suf[size-1] = 1
    for i := 1 ; i < size; i++ {
        pre[i] = pre[i-1] * nums[i-1]
    }
    for i := size - 2; i >= 0; i-- {
        suf[i] = suf[i+1] * nums[i+1]
    }
    for i := 0; i < size; i++ {
        res[i] = pre[i] * suf[i]
    }
    return res
}
