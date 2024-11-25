// https://leetcode.com/problems/removing-stars-from-a-string/
package stack
func removeStars(s string) string {
    n := len(s)

    stk := make([]byte, 0, n)
    top := 0

    for i := 0; i < n; i++ {
        if s[i] == '*' && top > 0 {
            stk = stk[:top - 1]
            top--
        } else {
            stk = append(stk, s[i])
            top++
        }
    }
    return string(stk)
}
