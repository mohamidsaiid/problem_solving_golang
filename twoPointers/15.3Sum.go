// https://leetcode.com/problems/3sum/
package twoPointers
import "slices"

func threeSum(nums []int) [][]int {
    n := len(nums)
    slices.Sort(nums)
    res := make([][]int, 0, n)

    for i := 0; i < n ; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        l, r := i + 1, n-1

        for l < r {
            tmp := nums[i] + nums[l] + nums[r]
            
            if tmp > 0 {
                r--
            } else if tmp < 0 {
                l++
            } else {
                triplet := []int{ nums[i], nums[l], nums[r] }
                res = append(res, triplet)
                // to avoid duplicates
                for l < r && nums[l] == nums[l+1] {
                    l++
                }
                for l < r && nums[r] == nums[r-1] {
                    r--
                }

                l++
                r--
            }
        }
    }
    return res
}
