// https://leetcode.com/problems/reverse-vowels-of-a-string/
package arrays
func vowel(c rune) bool {
    switch c{
    case 'a', 'e', 'o', 'i', 'u', 'A', 'E', 'O', 'I', 'U':
        return true
    default:
        return false
    }
}
func reverseVowels(s string) string {
    
    size := len(s)

    res := []rune(s)
    l, r := 0, size - 1

    for l < r {
        for l < r && !vowel(res[l]) {
            l++
        }
        for l < r && !vowel(res[r]) {
            r--
        }
        res[l], res[r] = res[r], res[l]
        r--
        l++
    }
    return string(res)
}
