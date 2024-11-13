// https://leetcode.com/problems/minimum-absolute-difference/
package slidingWindow
import(
    "slices"
    "math"
)
func minimumAbsDifference(arr []int) [][]int {
    n := len(arr)

    slices.Sort(arr)

    minn := math.MaxInt32
    for i := 0; i < n-1; i++ {
        minn = min(minn, arr[i+1] - arr[i])
    }

    res := make([][]int, 0, n)
    for i := 0; i < n-1; i++ {
        if arr[i+1] - arr[i] == minn {
            res = append(res, []int{arr[i], arr[i+1]})
        }
    }
    return res
}
