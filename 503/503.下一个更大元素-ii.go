/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 */

// @lc code=start
package main

func nextGreaterElements(nums []int) []int {
	// 单调栈
	stack := []int{}
	sl := 0
	l := len(nums)
	ans := make([]int, l)
	copy(ans, nums)
	for i := 0; i < 2*l; i++ {
		j := i % l
		if ans[j] == nums[j] {
			ans[j] = -1
		}
		for sl > 0 && nums[stack[sl-1]] < nums[j] {
			ans[stack[sl-1]] = nums[j]
			stack = stack[:sl-1]
			sl--
		}
		stack = append(stack, j)
		sl++
	}
	return ans
}

// @lc code=end
