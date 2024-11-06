// https://leetcode.com/problems/longest-mountain-in-array/description/
package twoPointers

func longestMountain(arr []int) int {
    n := len(arr)
    maxx := 0

    for i := 1; i < n - 1; i++ {

        if arr[i] > arr[i-1] && arr[i] > arr[i+1] {

            l, r := i, i
            
            for l > 0 && arr[l] > arr[l-1]  {
                l--
            }

            for r < n-1 && arr[r] > arr[r+1] {
                r++
            }
            if r != i && l != i {
                maxx = max(r-l+1, maxx)
            } 

        }

    }

    return maxx
}
