// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

package arrays

import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {

    kids := len(candies)

    maxCandies := slices.Max(candies)
    
    res := make([]bool, kids)
    for i := 0; i < kids; i++ {
        if candies[i] + extraCandies >= maxCandies {
            res[i] = true
        }
    }
    return res
}
