// https://leetcode.com/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {
    mp := make(map[int]struct{})
    for _, v := range nums {
        if _, ok := mp[v]; ok{
            return true
        }
        mp[v] = struct{}{}
    }
    return false
}
