// https://leetcode.com/problems/find-greatest-common-divisor-of-array/
package arrays
import "slices"
func ggcd (a, b int) int {
    if a == 0 {
        return b
    }
    return ggcd(b % a, a)
}
func findGCD(nums []int) int {
    maxx, minn := slices.Max(nums), slices.Min(nums)
    return ggcd(maxx, minn)
}
