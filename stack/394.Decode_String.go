package stack
import "strings"
func decodeString(s string) string {
    isDigit := func(d byte) bool {
        if d >= '0' && d <= '9' {
            return true
        }
        return false
    }
    strReverse := func(s string) string {
        var rev strings.Builder
        len := len(s)
        for i := len-1; i >= 0; i-- {
            rev.WriteByte(s[i])
        }
        return rev.String()
    }

    size := len(s)

    stk := make([]byte, 0, size)

    for i := 0; i < size; i++ {
        if s[i] != ']' {
            stk = append(stk, s[i])
            continue
        }
        var str, number strings.Builder
        
        for len(stk) > 0 && stk[len(stk) - 1] != '[' {
            str.WriteByte(stk[len(stk) - 1])
            stk = stk[:len(stk) - 1]
        }

        stk = stk[:len(stk) - 1] // to remove the [ side
        
        for len(stk) > 0 && isDigit(stk[len(stk) - 1]) {
            number.WriteByte(stk[len(stk) - 1])
            stk = stk[:len(stk) - 1]
        }
        resStr, resNumber := strReverse(str.String()), strReverse(number.String())
        mul, _ := strconv.Atoi(resNumber)
        res := strings.Repeat(resStr, mul)
        
        stk = append(stk, []byte(res)...)
    }
    
    return string(stk)
}
