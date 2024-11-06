package main

import "fmt"

func longestMountain(arr []int) int {
    n := len(arr)
    maxx := 0

    for i := 1; i < n - 1; i++ {

        if arr[i] > arr[i-1] && arr[i] > arr[i+1] {

            l, r := i, i
            fmt.Println(arr[i])
            
            for l > 0 && arr[l] > arr[l-1]  {
                l--
            }

            for r < n && arr[r] > arr[r+1] {
                r++
            }
            fmt.Println(l, r)
            if r != i && l != i {
                maxx = max(r-l+1, maxx)
            } 

        }

    }

    fmt.Println()
    return maxx
}

func main() {
    arr := []int{2,1,4,7,3,2,5}
    fmt.Println(longestMountain(arr))
}
