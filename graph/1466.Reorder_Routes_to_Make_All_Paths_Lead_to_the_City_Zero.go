package graph

func minReorder(n int, connections [][]int) int {
    neighbours := make([][]int, n)
    edges := make(map[int][]int)

    for _, edge := range connections {
        from, to := edge[0], edge[1]
        neighbours[from] = append(neighbours[from], to)
        neighbours[to] = append(neighbours[to], from)
        edges[from] = append(edges[from], to)
    }
    visited := make([]bool, n)
    count := 0
    find := func(x, y int) bool {
        v, ok := edges[x]
        if ok {
            for _, val := range v {
                if val == y {
                    return true
                }
            }
        }        
        return false
    }

    var dfs func(city int)
    dfs = func(city int) {
        
        for _, neighbour := range neighbours[city] {
            if visited[neighbour] {
                continue
            }
            if !find(neighbour, city){
                count++
            }
            visited[neighbour] = true
            dfs(neighbour)
        }
    }
    visited[0] = true
    dfs(0)
    return count 
}
