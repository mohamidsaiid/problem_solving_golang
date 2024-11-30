package graph

type coordinates struct {
    x int
    y int
}
func abs (x int) int {
    if x < 0 {
        return x * -1
    }
    return x
}
func nearestExit(maze [][]byte, entrance []int) int {
    dx := []int{1, -1, 0, 0}
    dy := []int{0, 0, 1, -1}
    n, m := len(maze), len(maze[0])
    cameFrom := make(map[coordinates]coordinates)

    validEdges := make(map[coordinates]int)

    for i := 0; i < m; i++ {
        if maze[0][i] != '+' {
            validEdges[coordinates{0, i}] = entrance[0] + abs(entrance[1] - i)
        }
    }

    for i := 0; i < n; i++ {
        if maze[i][0] != '+' {
            validEdges[coordinates{i, 0}] = abs(entrance[0] - i) + entrance[1] 
        }
    }

    for i := 0; i < n; i++ {
        if maze[i][m-1] != '+' {
            validEdges[coordinates{i, m - 1}] = abs(entrance[0] - i) + abs(entrance[1] - m + 1)        
        }
    }

    for i := 0; i < m; i++ {
        if maze[n - 1][i] != '+' {

            validEdges[coordinates{n - 1, i}] = abs(entrance[1] - i) + abs(entrance[0] - n + 1)    
        }
    }

    if _, ok := validEdges[coordinates{entrance[0], entrance[1]}]; ok {
        delete(validEdges, coordinates{entrance[0], entrance[1]})
    }
    if len(validEdges) == 0 {
        return -1
    }

    valid := func(x, y int) bool {
        c := coordinates{x, y}
        if x < 0 || y < 0 || x >= n || y >= m || maze[x][y] == '+' {
            return false
        }
        if _, ok := cameFrom[c]; ok {
            return false
        }

        return true
    }

    queue := make([]coordinates, 0, 100)
    queue = append(queue, coordinates{entrance[0], entrance[1]})
    cameFrom[coordinates{entrance[0], entrance[1]}] = coordinates{}

    for len(queue) != 0 {
        current := queue[0]
        queue = queue[1:]
        for i := 0; i < 4; i++ {
            nextx, nexty := current.x + dx[i], current.y + dy[i]
            if !valid(nextx, nexty) {
                continue
            }
            next := coordinates{nextx, nexty}
            cameFrom[next] = current
            queue = append(queue, next)
        }
    }
    fmt.Println(cameFrom)
    if len(cameFrom) == 0 {
        return -1
    }

    path := func(current, start coordinates) int {
        p := 0
        for current != start {
            p++
            current = cameFrom[current]
        }
        return p
    }

    minn := math.MaxInt
    for edge, _ := range validEdges {
        if _, ok := cameFrom[edge]; !ok {
            continue
        }
        e := coordinates{entrance[0], entrance[1]}
        tmp := path(edge, e)
        //fmt.Println(tmp)
        minn = min(minn, tmp)
    }
    if minn == math.MaxInt {
        return -1
    }
    fmt.Println(validEdges)
    return minn
}
