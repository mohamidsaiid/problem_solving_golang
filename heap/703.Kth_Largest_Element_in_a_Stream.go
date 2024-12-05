package heap
import "math"
type KthLargest struct {
    Data []int
    Times int
}


func Constructor(k int, nums []int) KthLargest {
    kthLargest := KthLargest{[]int{math.MinInt}, k}
    for _, val := range nums {
        kthLargest.Add(val)
    }
    return kthLargest
}


func (this *KthLargest) Add(val int) int {
    if len(this.Data) < this.Times+1 || this.top() < val {
        this.push(val)

        if len(this.Data) > this.Times +1{
            this.pop()
        }
    }
    return this.top()
}

func (this KthLargest) top() int {
    if len(this.Data) > 1 {
        return this.Data[1]
    }
    return math.MinInt
}


func (h *KthLargest) pop() int {
    heapSize := len(h.Data)
    if heapSize == 1 {
        return -1
    }
    if heapSize == 2 {
        top := h.Data[1]
        h.Data = h.Data[:1]
        return top
    }
    
    res := h.Data[1]
    h.Data[1] = h.Data[heapSize - 1]
    heapSize--
    h.Data = h.Data[:heapSize]

    i := 1
    for i * 2 < heapSize {
        if i*2+1 < heapSize && h.Data[i*2+1] < h.Data[i*2] && h.Data[i] > h.Data[i*2+1] {
            h.Data[i], h.Data[i*2+1] = h.Data[i*2+1], h.Data[i]
            i = i*2+1
        } else if h.Data[i*2] < h.Data[i] {
            h.Data[i], h.Data[i*2] = h.Data[i*2], h.Data[i]
            i = i*2
        } else {
            break
        }
    }
    return res
}



func (this *KthLargest) push(val int) {
    this.Data = append(this.Data, val)
    size := len(this.Data)
    i := size - 1
    for i/2 > 0 && this.Data[i] < this.Data[i/2] {
        this.Data[i], this.Data[i/2] = this.Data[i/2], this.Data[i]
        i /= 2
    }
}
/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
