package graph
func findCircleNum(isConnected [][]int) int {
    noOfCities := len(isConnected)
    adjList := make([][]int, noOfCities)
    reachedCities := make([]bool, noOfCities)

    for i := 0; i < noOfCities; i++ {
        for j := 0; j < noOfCities; j++ {
            if i == j {
                continue
            }
            if isConnected[i][j] == 1 {
                adjList[i] = append(adjList[i], j)
                //adjList[j] = append(adjList[j], i)
            }
        }
    }

    var dfs func(city int)
    dfs = func(city int) {
        reachedCities[city] = true
        for _, neighbour := range adjList[city] {
            if !reachedCities[neighbour] {
                dfs(neighbour)
            }
        }
    }
    noOfProvinces := 0
    for i := 0; i < noOfCities; i++ {
        if !reachedCities[i] {
            dfs(i)
            noOfProvinces++
        }
    }
    return noOfProvinces
}
