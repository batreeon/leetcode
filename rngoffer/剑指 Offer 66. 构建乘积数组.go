func constructArr(a []int) []int {
    l := len(a)
    if l == 0 {return []int{}}
    left := make([]int,l)
    left[0] = a[0]
    right := make([]int,l)
    right[l-1] = a[l-1]
    for i := 1 ; i < l-1 ; i++ {
        left[i],right[l-1-i] = left[i-1]*a[i],right[l-i]*a[l-1-i]
    }
    result := make([]int,l)
    result[0],result[l-1] = right[1],left[l-2]
    for i := 1 ; i < l-1 ; i++ {
        result[i] = left[i-1]*right[i+1]
    }
    return result
}