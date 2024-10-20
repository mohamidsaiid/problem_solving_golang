// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/
func smallerNumbersThanCurrent(nums []int) []int {
    n := len(nums)
    sortedArr := make([]int, n)
    copy(sortedArr, nums)
    slices.Sort(sortedArr)

    toIndex := make(map[int]int)
    for idx, v := range sortedArr {
        if _, ok := toIndex[v]; !ok {
            toIndex[v] = idx
        }
    }
    
    res := make([]int, n)
    for i, v := range nums {
        res[i] = toIndex[v]
    }
    return res
}
