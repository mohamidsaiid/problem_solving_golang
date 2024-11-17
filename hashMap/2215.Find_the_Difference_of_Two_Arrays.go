// https://leetcode.com/problems/find-the-difference-of-two-arrays/
package hashMap
func findDifference(nums1 []int, nums2 []int) [][]int {
    mp1, mp2 := make(map[int]bool), make(map[int]bool)
    for _, num := range nums1 {
        mp1[num] = true
    }
    for _, num := range nums2 {
        mp2[num] = true
    }
    res := make([][]int, 2)
    for num, _ := range mp1 {
        if _, ok := mp2[num]; !ok {
            res[0] = append(res[0], num)
            delete(mp2, num)
        }
    }
    for num, _ := range mp2 {
        if _, ok := mp1[num]; !ok {
            res[1] = append(res[1], num)
            delete(mp1, num)
        }
    }
    return res
}
