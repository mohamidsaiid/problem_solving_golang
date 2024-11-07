// https://leetcode.com/problems/contains-duplicate-ii/
package slidingWindow

func containsNearbyDuplicate(nums []int, k int) bool {
    size := len(nums)

    visited := make(map[int]struct{})
    for i := 0; i < size; i++ {
        if _, ok := visited[nums[i]]; ok {
            return true
        }
        visited[nums[i]] = struct{}{}
        if len(visited) > k {
            delete(visited, nums[i-k])
        }
    }
    return false
}
