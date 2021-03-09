/*
给你一个整数数组 nums​​​ 和一个整数 k​​​​​ 。区间 [left, right]（left <= right）的 异或结果 是对下标位于 left 和 right（包括 left 和 right ）之间所有元素进行 XOR 运算的结果：nums[left] XOR nums[left+1] XOR ... XOR nums[right] 。

返回数组中 要更改的最小元素数 ，以使所有长度为 k 的区间异或结果等于零。
011	100 111
*/
package main
import "math"


func minChanges(nums []int, k int) int {
	mintime := math.MaxInt64
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	n := len(nums)
	numsCopy := make([]int,n)
	for change := k-1 ; change >= 0 ; change-- {
		copy(numsCopy,nums)
		pre := 0
		r := k-1
		for i := 0 ; i <= r ; i++ {
			if i != change {
				pre ^= numsCopy[i]
			}
		}
		times := 0
		if pre != numsCopy[change] {
			numsCopy[change] = pre
			times++
		}
		r++
		for r < n {
			if numsCopy[r] != numsCopy[r-k] {
				numsCopy[r] = numsCopy[r-k]
				times++
			}
			r++
		}
		mintime = min(mintime,times)
	}
	r := k
	for r < n {
		pre := 0
		for i := r-k+1 ; i < r ; i++ {
			pre ^= numsCopy[i]
		}
		times := 0
		if pre != numsCopy[r] {
			numsCopy[r] = pre
			times++
		}
		for i := 0 ; i < n ; i++ {
			if i < r-k+1 {
				if nums
			}
		}
	}
	return mintime
}