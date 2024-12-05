package heap
import (
    "sort"
    "container/heap"
    "fmt"
)

type element struct {
    n2 int
    n1 int
}
type arr []element
func (a arr) Len() int { return len(a) }
func (a arr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a arr) Less(i, j int) bool { return a[i].n2 > a[j].n2 }

func maxScore(nums1 []int, nums2 []int, k int) int64 {
    array := arr{}
    lengthOfNums := len(nums1)
    for i := 0; i < lengthOfNums; i++ {
        array = append(array, element{nums2[i], nums1[i]})
    }
    sort.Sort(array)
    h := new(IntHeap)
    heap.Init(h)
    maxx := 0
    totSum := 0
    for _, a := range array {
        totSum += a.n1
        heap.Push(h, a.n1)
        if h.Len() > k {
            n1Pop := heap.Pop(h).(int)
            totSum -= n1Pop
        }
        if h.Len() == k {
            maxx = max(maxx, totSum * a.n2)
        }
    }
    //maxx = max(maxx, totSum * array[lengthOfNums-1].n2)
    fmt.Println(array)
    fmt.Println(totSum, maxx)
    return int64(maxx)
}
