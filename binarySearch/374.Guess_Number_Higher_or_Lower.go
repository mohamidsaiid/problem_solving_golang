package binarySearch
/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 */

func guess(num int) int;
func guessNumber(n int) int {
    s, e := 1, n
    for s <= e {
        mid := s + (e - s) / 2
        tmp := guess(mid)
        if tmp == 0 {
            return mid 
        } else if tmp == -1 {
            e = mid-1
        } else {
            s = mid+1
        }
    }
    return n
}
