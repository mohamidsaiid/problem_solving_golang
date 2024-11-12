// https://leetcode.com/problems/merge-strings-alternately/
package arrays
import "strings"
func mergeAlternately(word1 string, word2 string) string {
    var res strings.Builder
    word1Length, word2Length := len(word1), len(word2)

    for i := 0; i < word1Length + word2Length; i++ {
        if i < word1Length {
            res.WriteByte(word1[i])
        }
        if i < word2Length {
            res.WriteByte(word2[i])
        }
    }
    return res.String()
}
