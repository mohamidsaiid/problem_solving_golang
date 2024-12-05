package heap
import (
    "container/heap"
)
type MinHeap []int
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(val any) {
    *h = append(*h, val.(int))
}
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h MinHeap) top() int{
    return h[0]
}
func totalCost(costs []int, k int, candidates int) int64 {
    headWorkers, tailWorkers := new(MinHeap), new(MinHeap)
    costsLength := len(costs)

    for i := 0; i < candidates; i++ {
        heap.Push(headWorkers, costs[i])
    }
    for i := max(candidates, costsLength - candidates); i < costsLength; i++ {
        heap.Push(tailWorkers, costs[i])
    }

    ans := 0
    nextHead, nextTail := candidates, costsLength - 1 - candidates

    for i := 0; i < k; i++ {
        if tailWorkers.Len() == 0 || headWorkers.Len() != 0 && headWorkers.top() <= tailWorkers.top() {
            ans += heap.Pop(headWorkers).(int)

            if nextHead <= nextTail {
                heap.Push(headWorkers, costs[nextHead])
                nextHead++
            }
        } else {
            ans += heap.Pop(tailWorkers).(int)

            if nextHead <= nextTail {
                heap.Push(tailWorkers, costs[nextTail])
                nextTail--
            }
        }
    
    }
    return int64(ans)
}
