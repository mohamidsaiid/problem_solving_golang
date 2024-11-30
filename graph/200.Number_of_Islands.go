// https://leetcode.com/problems/number-of-islands/
package graph
// bfs solution
func numIslands(grid [][]byte) int {
    n, m := len(grid), len(grid[0])
    di, dj := []int{1, -1, 0, 0}, []int{0, 0, 1, -1}
    islands := 0

    valid := func (i, j int) bool {
        if i < n && i >= 0 && j < m && j >= 0 && grid[i][j] == '1'{
            return true
        }
        return false
    }
    var bfs func(i, j int) 
    bfs = func(i, j int) {
        q := [][]int{{i, j}}
        grid[i][j] = '0'

        for len(q) != 0 {
            cell := q[0]
            q = q[1:]
            
            for d := 0; d < 4; d++ {
                nextI := cell[0] + di[d]
                nextJ := cell[1] + dj[d]
                
                if valid(nextI, nextJ) {
                    q = append(q, []int{nextI, nextJ})
                    grid[nextI][nextJ] = '0' // visited
                }
            }
        }
    }

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] != '1' {
                continue
            }
            bfs(i, j)
            islands++
        }
    }
    return islands
}
