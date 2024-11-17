// https://leetcode.com/problems/unique-number-of-occurrences/
package hashMap
func uniqueOccurrences(arr []int) bool {
    freq := make([]int, 2001)
    for _, n := range arr {
        if n > 0 {
            freq[n-1]++
        } else {
            freq[2000+n]++
        }
    }
    cntUnique := 0
    unique := make(map[int]struct{})
    for i := 0; i < 2001; i++ {
        if freq[i] != 0 {
            cntUnique++
            unique[freq[i]] = struct{}{}
        }
    }
    if cntUnique > len(unique) {
        return false
    }
    return true
}
