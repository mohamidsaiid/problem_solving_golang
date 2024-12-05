package heap
import (
    "container/heap"
)
type IntHeap []int
func(h IntHeap) Less(i, j int) bool {return h[i] < h[j]}
func(h IntHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func(h IntHeap) Len() int {return len(h)}
func(h *IntHeap) Push(x any) {
    *h = append(*h, x.(int))
}
func(h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0:n-1]
    return x
}

type SmallestInfiniteSet struct {
    Data *IntHeap
    d []bool
}


func nConstructor() SmallestInfiniteSet {
    d := make([]bool, 1001)
    data := new(IntHeap)
    for i := 0; i < 1001; i++ {
        d[i] = true
        data.Push(i)
    }
    heap.Init(data)
    heap.Pop(data)
    return SmallestInfiniteSet{data, d}
}


func (this *SmallestInfiniteSet) PopSmallest() int {
    ele := heap.Pop(this.Data)
    this.d[ele.(int)] = false
    return ele.(int)
}


func (this *SmallestInfiniteSet) AddBack(num int)  {
    if this.d[num] {
        return
    }
    heap.Push(this.Data, num)
    this.d[num] = true
}


/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
