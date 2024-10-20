// https://leetcode.com/problems/contains-duplicate-ii/
func containsNearbyDuplicate(nums []int, k int) bool {
    mpToIndex := make(map[int]int)

    for idx, num := range nums {
        if val, ok := mpToIndex[num]; ok && idx - val <= k {
            return true
        }
        mpToIndex[num] = idx
    }
    return false
}
