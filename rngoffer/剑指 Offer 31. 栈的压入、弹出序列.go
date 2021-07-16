func validateStackSequences(pushed []int, popped []int) bool {
    // 模拟
    stack := []int{}
    pushIdx := 0
    for _,popnum := range popped {
        if len(stack) == 0 {
            stack = append(stack,pushed[pushIdx])
            pushIdx++
        }
        for pushIdx < len(pushed) && stack[len(stack)-1] != popnum {
            stack = append(stack,pushed[pushIdx])
            pushIdx++
        }
        if stack[len(stack)-1] != popnum {
            return false
        }else{
            stack = stack[:len(stack)-1]
        }
    }
    return true
}