package backtrack
func combinationSum3(k int, n int) [][]int {
    res := make([][]int, 0, 10)
    var backtrack func(idx, val int, subset []int)
    backtrack = func(idx, val int, subset []int) {
        if len(subset) > k || val < 0{
            return
        }
        if len(subset) == k && val == 0 {
            tmp := make([]int, k)
            copy(tmp, subset)
            res = append(res, tmp)
            return
        }
        for i := idx; i <= 9; i++ {
            subset = append(subset, i)
            backtrack(i+1, val - i, subset)
            subset = subset[:len(subset)-1]
        }
    }
    backtrack(1, n, []int{}, )
    return res
}
