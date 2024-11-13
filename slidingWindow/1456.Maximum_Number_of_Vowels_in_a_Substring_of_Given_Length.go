// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/

package slidingWindow

func checkVowel(c byte) bool {
    return c == 'a' || c == 'e' || c == 'o' || c == 'i' || c == 'u'
}

func maxVowels(s string, k int) int {
    cntVowels := 0
    for i := 0; i < k; i++ {
        if checkVowel(s[i]) {
            cntVowels++
        }
    }
    maxx := cntVowels
    l, r := 0, k
    for r < len(s) {
        if checkVowel(s[l]) {
            cntVowels--
        } 
        if checkVowel(s[r]) {
            cntVowels++
        }
        r++
        l++
        maxx = max(maxx, cntVowels)
    }
    return maxx
}
