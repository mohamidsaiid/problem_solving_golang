// https://leetcode.com/problems/reverse-words-in-a-string/

package arrays

import "strings"

func reverseWords(s string) string {
    words := strings.Fields(s)
    
    res := ""    
    for i := len(words) - 1; i > 0; i-- {
        res += words[i] + " "
    }
    res += words[0] 
    return res
}
