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
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}
	nums = []int{}
	for k := range m {
		nums = append(nums, k)
	}
	sort.Ints(nums)

	result := 0
	for len(nums) != 1 {
		l := len(nums)
		if m[nums[0]]*(nums[1]-nums[0]) < m[nums[l-1]]*(nums[l-1]-nums[l-2]) {
			result += m[nums[0]]*(nums[1]-nums[0])
			m[nums[1]] += m[nums[0]]
			nums = nums[1:]
		}else{
			result += m[nums[l-1]]*(nums[l-1]-nums[l-2])
			m[nums[l-2]] += m[nums[l-1]]
			nums = nums[:l-1]
		}
	}
	return result
}
// @lc code=end

