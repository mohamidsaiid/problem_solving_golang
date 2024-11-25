// https://leetcode.com/problems/asteroid-collision/
package stack
func asteroidCollision(asteroids []int) []int {
    n := len(asteroids)
    stk := make([]int, 0, n)
    topPtr := 0
    
    abs := func(x int) int {
        if x < 0 {
            return x * -1
        }
        return x
    }

    for i := 0; i < n; i++{
        nxt := asteroids[i]

        if topPtr == 0 {
            stk = append(stk, nxt)
            topPtr++
            continue
        }
        lst := stk[topPtr - 1]

        if lst < 0 || nxt > 0 {
            stk = append(stk, nxt)
            topPtr++
            continue
        }
        
        absNext := abs(nxt)
        absLast := abs(lst)

        if (absLast > absNext) {
            continue
        } else if (absNext > absLast) {
            stk = stk[:topPtr - 1]
            topPtr--
            i--
        } else {
            stk = stk[:topPtr - 1]
            topPtr--
        }
    }
    return stk
}
