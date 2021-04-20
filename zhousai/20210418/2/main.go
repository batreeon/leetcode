func maxIceCream(costs []int, coins int) int {
    sort.Ints(costs)
    i := 0
    result := 0
    for ; i < len(costs) ; i++ {
        coins -= costs[i]
        if coins < 0 {
            break
        }
        result++
    }
    return result
}