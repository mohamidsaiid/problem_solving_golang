package binarySearch
import "slices"
func successfulPairs(spells []int, potions []int, success int64) []int {
    slices.Sort(potions)
    spellsSize := len(spells)
    res := make([]int, 0, spellsSize)
    for i := 0; i < spellsSize; i++ {
        l, r := 0, len(potions) - 1

        for l <= r {
            mid := l + (r - l) / 2
            if int64(spells[i] * potions[mid]) < success {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
        res = append(res, len(potions) - l)
    }
    return res
}
