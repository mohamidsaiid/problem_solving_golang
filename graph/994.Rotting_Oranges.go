package graph

func orangesRotting(grid [][]int) int {
    q := [][]int{}
    time := 0
    n, m := len(grid), len(grid[0])
    fresh := 0
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == 1 {
                fresh++
            }
            if grid[i][j] == 2 {
                q = append(q, []int{i, j})
            }
        }
    }

    adjx := []int{1, -1, 0, 0}
    adjy := []int{0, 0, 1, -1}
    for len(q) != 0 {
        ln := len(q)
        for i := 0; i < ln; i++ {
            x, y := q[0][0], q[0][1]
            
            for j := 0; j < 4; j++ {
                nx, ny := x + adjx[j], y + adjy[j]
                if nx >= n || ny >= m || nx < 0 || ny < 0 || grid[nx][ny] != 1 {
                    continue
                }
                q = append(q, []int{nx, ny})
                grid[nx][ny] = 2
                fresh--
            }
            q = q[1:]
        }
        time++
    }
    if fresh == 0 {
        if time == 0 {
            return 0
        }
        return time-1
    }
    return -1
}
