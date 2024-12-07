package binarySearch
func findPeakElement(nums []int) int {
    n := len(nums)
    if n == 1 {
        return 0
    } else if nums[0] > nums[1] {
        return 0
    } else if nums[n-1] > nums[n-2] {
        return n-1
    }

    l, r := 1, n - 2
    
    for l <= r {
        mid := l + (r-l) / 2
        if nums[mid] > nums[mid-1] && nums[mid + 1] < nums[mid] {
            return mid
        } else if  nums[mid] > nums[mid - 1] {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return l
}
