// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
package twoPointers

func maxProfit(prices []int) int {
    n := len(prices)

    l, r := 0, 1
    maxx := 0
    for r < n {
        prof := prices[r] - prices[l]
        if prof <= 0{
            l = r
        }else {
            maxx = max(maxx, prof)
        }
        r++
    }
    return maxx
}
