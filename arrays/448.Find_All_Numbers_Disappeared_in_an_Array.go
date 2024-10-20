// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
func findDisappearedNumbers(nums []int) []int {
    n := len(nums)
    cntSet := make([]bool, n+1)
    for _, val := range nums {
        cntSet[val-1] = true
    }
    res := make([]int, 0, n)
    for i := 0; i < n; i++ {
        if !cntSet[i] {
            res = append(res, i+1)
        }
    }
    return res
}
