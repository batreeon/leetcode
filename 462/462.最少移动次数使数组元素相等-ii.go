/*
 * @lc app=leetcode.cn id=462 lang=golang
 *
 * [462] 最少移动次数使数组元素相等 II
 */

// @lc code=start
package main
import (
	"sort"
)
func minMoves2(nums []int) int {
	sort.Ints(nums)
	t := nums[len(nums)/2]
	result := 0
	for _, num := range nums {
		result += abs(num - t)
	}
	return result
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
// @lc code=end

