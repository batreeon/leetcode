package main
/*
给你三个正整数 n、index 和 maxSum 。你需要构造一个同时满足下述所有条件的数组 nums（下标 从 0 开始 计数）：

nums.length == n
nums[i] 是 正整数 ，其中 0 <= i < n
abs(nums[i] - nums[i+1]) <= 1 ，其中 0 <= i < n-1
nums 中所有元素之和不超过 maxSum
nums[index] 的值被 最大化
返回你所构造的数组中的 nums[index] 。

注意：abs(x) 等于 x 的前提是 x >= 0 ；否则，abs(x) 等于 -x 。
*/
func maxValue(n int, index int, maxSum int) int {
	res := make([]int,n)
	for i := 0 ; i < n ; i++ {
		res[i] = 1
	}
	res[index]++
	maxsum := n+1
	if maxsum > maxSum {
		return 1
	}
	addNum := 1
	left,right := index,n-1-index
	min := func(x,y int) int {
		if x > y {
			return y
		}
		return x
	}
	for {
		res[index]++
		maxsum++
		maxsum += min(left,addNum)
		maxsum += min(right,addNum)
		addNum++
        if maxsum > maxSum {
			break
        }
        if addNum >= left && addNum >= right {
            nn := (maxSum-maxsum)/(1+left+right)
            res[index] += nn
            // maxsum += nn * left
            // maxsum += nn * right
            return res[index]
        }
		
	}
	indexMax := res[index]-1
	return indexMax
}