package heap
import (
    "container/heap"
    "fmt"
    "math"
)
type posHeap [][]int
func (p posHeap) Less(i, j int) bool {
    xi, yi, xj, yj := p[i][0], p[i][1], p[j][0], p[j][1]
    disti, distj := math.Sqrt(float64(xi*xi) + float64(yi*yi)), math.Sqrt(float64(xj*xj) + float64(yj*yj))
    return disti < distj
}
func (p posHeap) Len() int {return len(p)}
func (p posHeap) Swap(i, j int) {p[i], p[j] = p[j], p[i]}
func (p *posHeap) Push(pos any) {
    *p = append(*p, pos.([]int))
}
func (p *posHeap) Pop() any {
    old := *p
    n := len(old)
    x := old[n-1]
    *p = old[0:n-1]
    return x
}
func kClosest(points [][]int, k int) [][]int {
    
    x := posHeap(points)
    heap.Init(&x)
    fmt.Println(x)
    res := [][]int{}
    for i := 0; i < k; i++ {
        res = append(res, heap.Pop(&x).([]int))
    }
    fmt.Println(res)
    return res
}
