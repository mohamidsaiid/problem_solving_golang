package heap
import (
    "container/heap"
    "fmt"
)
type Heap []int
func (h Heap) Less(i, j int) bool {return h[i] < h[j]}
func (h Heap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h Heap) Len() int {return len(h)}
func (h *Heap) Push(x any) {
    *h = append(*h, x.(int))
}
func (h *Heap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0:n-1]
    return x
}

func findKthLargest(nums []int, k int) int {
    h := new(Heap)
    heap.Init(h)
    for _, val := range nums {
        if h.Len() < k || (*h)[0] < val {
            heap.Push(h, val)
            if h.Len() > k {
                heap.Pop(h)
            }
        }
    }
    fmt.Println(h)
    return (*h)[0]
}

