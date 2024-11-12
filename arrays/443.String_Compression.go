// https://leetcode.com/problems/string-compression/
package arrays

import ( 
    "fmt"
    "strconv"
)

func compress(chars []byte) int {
    size := len(chars)
    i := 0
    resStr := ""
    for i < size {
        resStr += string(chars[i])
        j := i
        for i < size-1 && chars[i] == chars[i+1] {
            i++
        }
        i++
        fmt.Println(j, i, j - i)
        if (i - j) > 1 {
            resStr += strconv.Itoa(i - j)
        }
    }
    for j := 0; j < len(resStr); j++ {
        chars[j] = resStr[j]
    }
    return len(resStr)
}
