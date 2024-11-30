package graph
func canVisitAllRooms(rooms [][]int) bool {
    noOfRooms := len(rooms)
    reachedRooms := make([]bool, noOfRooms)
    
    var dfs func(roomKey int)

    dfs = func(roomKey int) {
        reachedRooms[roomKey] = true
        for _, key := range rooms[roomKey] {
            if !reachedRooms[key] {
                dfs(key)
            }
        }
    }

    dfs(0)
    reachedRooms[0] = true
    for _, k := range reachedRooms {
        if k == false {
            return false
        }
    }
    return true
}
