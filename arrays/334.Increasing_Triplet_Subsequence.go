// https://leetcode.com/problems/increasing-triplet-subsequence/
package arrays

import "math"

func increasingTriplet(nums []int) bool {
    num1, num2 := math.MaxInt32, math.MaxInt32
    for _, num := range nums {
        if num <= num1 {
            num1 = num
        } else if num <= num2 {
            num2 = num
        } else {
            return true
        }
    }
    return false
}
