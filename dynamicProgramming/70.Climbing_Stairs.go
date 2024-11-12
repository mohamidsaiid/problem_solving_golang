package DB
func climbStairs(n int) int {
    arr := make([]int, n+1)

    var dp func(x int) int 
    dp = func(x int) int{
        if arr[x] != 0 {
            return arr[x]
        }
        if x == 1 || x == 0 {
            return 1
        }
        arr[x] = dp(x-1) + dp(x-2)
        return arr[x]
    }
    return dp(n)
}
