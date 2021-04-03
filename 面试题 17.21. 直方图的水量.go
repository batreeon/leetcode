func trap(height []int) (ans int) {
    stack := []int{}
    min := func(x,y int) int {
        if x < y {
            return x
        }
        return y
    }
    result := 0
    for i,h := range height {
        for len(stack) > 0 && h > height[stack[len(stack)-1]] {
            water := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if len(stack) == 0 {
                // 右边没有更高的，存不住水，出现在右边起始处
                break
            }
            // 栈内的都是递减的，因此再往左取一个要么和上一个一样高，要么可作为左边界
            // 一样高，那么这次给result加上去的就是0，等到下一次遇到左边界再一次加上去
            // 遇到左边界就可以直接加了，见到height[top]，是因为把下面更低一层的已经加过了，可看作一个平台
            left := stack[len(stack)-1]
            wide := i - left - 1
            hheight := min(h,height[left]) - height[water]
            result += wide*hheight
        }
        stack = append(stack,i)
    }
    return result
}