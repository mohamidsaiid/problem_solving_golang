// https://leetcode.com/problems/is-subsequence/
package twoPointers
func isSubsequence(s string, t string) bool {
    i, j := 0, 0
    n, m := len(s), len(t)
    for i < n && j < m {
        if s[i] == t[j] {
            i++
        }
        j++
    }
    if i == n {
        return true
    }
    return false
}
