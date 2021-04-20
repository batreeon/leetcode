func getXORSum(arr1 []int, arr2 []int) int {
    xor1,xor2 := 0,0
    for _,arr := range arr1 {
        xor1 ^= arr
    }
    for _,arr := range arr2 {
        xor2 ^= arr
    }
    return xor1 & xor2
}