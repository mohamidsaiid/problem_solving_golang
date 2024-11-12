// https://leetcode.com/problems/can-place-flowers/
package arrays
func canPlaceFlowers(flowerbed []int, n int) bool {
    if n == 0 {
        return true
    }

    size := len(flowerbed)
    
    if size == 1 {
        if flowerbed[0] == 0 && n <= 1{
            return true
        } else {
            return false
        }
    }

    tot := 0
    if flowerbed[0] == 0 && flowerbed[1] == 0 {
        tot++
        flowerbed[0] = 1
    }
    
    if flowerbed[size - 1] == 0 && flowerbed[size - 2] == 0 {
        tot++
        flowerbed[size-1] = 1
    }

    for i := 1; i < size - 1; i++ {
        if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
            tot++
            flowerbed[i] = 1
        }
    }
    if tot >= n {
        return true
    }
    return false
}
