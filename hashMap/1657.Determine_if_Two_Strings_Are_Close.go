// https://leetcode.com/problems/determine-if-two-strings-are-close/
package hashMap
import "sort"
func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }

    solve := func(s string) []int {
        arr := make([]int, 26)
        for _, c := range s {
            arr[c - 'a']++
        }
        return arr
    }
    w1, w2 := solve(word1), solve(word2)

    for i := 0; i < 26; i++ {
        if (w1[i] > 0 && w2[i] == 0) || (w1[i] == 0 && w2[i] > 0) {
            return false
        }
    }
    
    sort.Ints(w1)
    sort.Ints(w2)
    for i := 0; i < 26; i++ {
        if w1[i] != w2[i] {
            return false
        }
    }
    return true
}
