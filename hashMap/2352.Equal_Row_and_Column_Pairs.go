// https://leetcode.com/problems/equal-row-and-column-pairs/
package hashMap
func equalPairs(grid [][]int) int {
    n := len(grid)
    rows := make(map[string]int)
    //cols := make(map[string]bool)
    
    for i := 0; i < n; i++ {
        r:= ""
        for j := 0; j < n; j++ {
            r += string(grid[i][j] + '0')
        }
        rows[r]++
    }

    cnt := 0
    for i := 0; i < n; i++ {
        s := ""
        for j := 0; j < n; j++ {
            s += string(grid[j][i] + '0')
        }
        if val, ok := rows[s]; ok {
            cnt += val
        }
    }
    return cnt
}
