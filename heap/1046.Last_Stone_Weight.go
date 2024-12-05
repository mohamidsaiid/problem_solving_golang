package heap

func lastStoneWeight(stones []int) int {
    if len(stones) == 1 {
        return stones[0]
    }

    rearrange := func(i int) {
        lenStones := len(stones)

        for i * 2 < lenStones {
            if i * 2 + 1 < lenStones && stones[i*2+1] > stones[i*2] && stones[i] < stones[i*2+1] {
                stones[i], stones[i*2+1] = stones[i*2+1], stones[i]
                i = i*2+1 
            } else if stones[i*2] > stones[i] {
                stones[i], stones[i*2] = stones[i*2], stones[i]
                i = i*2
            } else {
                break
            }
        }
    }

    heapify := func() {
        stones = append(stones, stones[0])
        stones[0] = math.MinInt
        current := (len(stones) - 1)/2
        for current > 0 {
            rearrange(current)
            current--
        }
    }
    pop := func() int {
        if len(stones) == 2 {
            tmp := stones[1] 
            stones = stones[:1]
            return tmp
        }
        res := stones[1]
        stones[1] = stones[len(stones) - 1]
        stones = stones[:len(stones) - 1]
        rearrange(1)
        return res
    }
    heapify()
    push := func(val int) {
        stones = append(stones, val)
        i := len(stones) - 1
        for i/2 > 0 && stones[i] > stones[i/2] {
            stones[i], stones[i/2] = stones[i/2], stones[i]
            i = i/2
        }
    }
    fmt.Println(stones)
    for len(stones) > 2 {
        y := pop()
        x := pop()
        fmt.Println(y, x)
        if y - x > 0 {
            push(y-x)

        }
        fmt.Println(stones)
    }
    fmt.Println(stones)
    if len(stones) == 1 {
        return 0
    }
    return stones[1]
}
