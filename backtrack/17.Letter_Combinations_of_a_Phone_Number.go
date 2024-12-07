package backtrack
import "strings"
func letterCombinations(digits string) []string {

    digitsSize := len(digits)
    if digitsSize == 0 {
        return []string{}
    }
    digitsMap := map[byte]string {
        '2' : "abc",
        '3' : "def",
        '4' : "ghi",
        '5' : "jkl",
        '6' : "mno",
        '7' : "pqrs",
        '8' : "tuv",
        '9' : "wxyz",
    }

    res := make([]string, 0, 10)

    var backtrack func(digitIdx int, perm string)
    backtrack = func (digitIdx int, perm string) {
        if digitIdx == digitsSize {
            tmp := strings.Clone(perm)
            res = append(res, tmp)
            return
        }
        s := digitsMap[digits[digitIdx]]
        for i := 0 ; i < len(s); i++ {
            perm += string(s[i])
            backtrack(digitIdx+1, perm)
            perm = perm[:len(perm) - 1]
        }
        
    }
    backtrack(0, "")
    return res
}
