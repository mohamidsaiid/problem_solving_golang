package binarySearch
import "slices"
import "math"
func minEatingSpeed(piles []int, h int) int {
    l, r := 1, slices.Max(piles)
    size := len(piles)
    minn := r
    for l <= r {
        mid := l + (r-l) / 2
        x := 0
        for i := 0; i < size; i++ {
            x += int(math.Ceil(float64(piles[i])/ float64(mid)))
        }
        if x <= h {
            minn = min(minn, mid)
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return minn
}
