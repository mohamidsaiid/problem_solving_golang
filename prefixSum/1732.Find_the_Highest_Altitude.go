// https://leetcode.com/problems/find-the-highest-altitude/
package prefixSum
func largestAltitude(gain []int) int {
    highest := 0
    alt := 0
    for i := 0; i < len(gain); i++ {
        alt = alt + gain[i]
        highest = max(highest, alt)
    }
    return highest
}
