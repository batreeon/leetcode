/*
 * @lc app=leetcode.cn id=1005 lang=golang
 *
 * [1005] K 次取反后最大化的数组和
 */

// @lc code=start
package main

import "sort"

func largestSumAfterKNegations(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	result := 0
	if nums[0] >= 0 {
		for _, num := range nums {
			result += num
		}
	}else{
		
	}
}
// @lc code=end

