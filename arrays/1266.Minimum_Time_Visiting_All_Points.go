package arrays
// https://leetcode.com/problems/minimum-time-visiting-all-points/
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func minTimeToVisitAllPoints(points [][]int) int {
    time := 0
    size := len(points)

    for i := 1; i < size ; i++ {
        startPoint := points[i-1]
        nextPoint := points[i]

        a := abs(startPoint[0] - nextPoint[0])
        b := abs(startPoint[1] - nextPoint[1])
        time += max(a,b)
    }
    return time
}
