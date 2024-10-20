// https://leetcode.com/problems/missing-number/
func missingNumber(nums []int) int {
    n := len(nums)
    sum, totalSum := 0, n*(n+1)/2 
    for _, v := range nums {
        sum += v
    }
    return totalSum - sum
}
