/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小操作次数使数组元素相等
 */

// @lc code=start
func minMoves(nums []int) int {
	// 给n-1个数加1，相当于给一个数减1，那么操作次数等于所有数字减小为最小数的操作次数
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	result := 0
	for _, num := range nums {
		result += num - min
	}

	return result
}
// @lc code=end

