package queue
func predictPartyVictory(senate string) string {
    senLen := len(senate)
    r, d := make([]int, 0, senLen), make([]int, 0, senLen)

    for i := 0; i < senLen; i++ {
        if senate[i] == 'R' {
            r = append(r, i)
        } else {
            d = append(d, i)
        }
    }

    for len(r) > 0 && len(d) > 0 {
        rTurn, dTurn := r[0], d[0]
        r, d = r[1:], d[1:]
        if rTurn < dTurn {
            r = append(r, rTurn+senLen)
        } else {
            d = append(d, dTurn+senLen)
        }
    }
    if len(r) == 0 {
        return "Dire"
    }
    return "Radiant"
}
