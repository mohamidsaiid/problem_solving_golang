package graph

type A struct {
    d string
    Val float64
}
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    res := make([]float64, len(queries))

    adjList := make(map[string][]A)
    for idx, eq := range equations {
        d1, d2 := A{eq[1], values[idx]}, A{eq[0], 1/values[idx]}
        adjList[eq[0]] = append(adjList[eq[0]], d1)
        adjList[eq[1]] = append(adjList[eq[1]], d2)
    }

    queryAns := 0.0
    visi := make(map[string]bool)
    var dfs func(s, e string, runningProduct float64) bool
    dfs = func(s, e string, runningProduct float64) bool {
        if s == e {
            queryAns = runningProduct
            return true
        }
        tmp := false
        visi[s] = true
        for _, v := range adjList[s] {
            dist, val := v.d, v.Val
            if visi[dist] {
                continue
            }


            tmp = dfs(dist, e, runningProduct * val)
            if tmp {
                break
            } 

        }
        visi[s] = false
        return tmp
    }

    for idx, q := range queries {
        if _, ok := adjList[q[0]]; !ok {
            res[idx] = -1.0
            continue
        }
        if _, ok := adjList[q[1]]; !ok {
            res[idx] = -1.0
            continue
        }
        queryAns = 1.0
        if dfs(q[0], q[1], 1) {
            res[idx] = queryAns
        } else {
            res[idx] = -1
        }
    }

    return res
}
