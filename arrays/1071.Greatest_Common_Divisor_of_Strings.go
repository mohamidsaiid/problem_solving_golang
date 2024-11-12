// https://leetcode.com/problems/greatest-common-divisor-of-strings/
package arrays
func gcd (a, b int) int {
    if a == 0 {
        return b
    }
    return gcd(b % a, a)
}

func gcdOfStrings(str1 string, str2 string) string {
    if str1 + str2 != str2 + str1 {
        return ""
    }

    gcdLen := gcd(len(str1), len(str2))
    return str1[:gcdLen]    
}


/* another brute force soultion
func gcdOfStrings(str1 string, str2 string) string {
    if str1 + str2 != str2 + str1 {
        return ""
    }
    len1, len2 := len(str1), len(str2)

    joinWords := func (str string, k int) string {
        ans := ""
        for i := 0; i < k; i++ {
            ans += str
        }
        return ans
    }

    valid := func (k int) bool {
        if len1 % k != 0 || len2 % k != 0 {
            return false
        }
        base := str1[:k]
        n1, n2 := len1 / k, len2 / k
        return str1 == joinWords(base, n1) && str2 == joinWords(base, n2)
    }


    for i := min(len1, len2); i >= 1; i-- {
        if valid(i) {
            return str1[:i]
        }
    } 
    return ""
}
*/
