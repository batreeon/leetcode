/*
 * @lc app=leetcode.cn id=1968 lang=golang
 *
 * [1968] 构造元素不等于两相邻元素平均值的数组
 */

// @lc code=start
package main

import (
	"sort"
)

func rearrangeArray(nums []int) []int {
	// 排序，前半部分中任一数小于后半部分的任一数，因此将这两部分交叉排在一起
	// 一定不会某一元素等于左右两元素均值的情况
	sort.Ints(nums)
	n := len(nums)
	m := (n+1)/2
	result := []int{}
	for i := 0 ; i < m ; i++ {
		result = append(result,nums[i])
		if i+m < n {
			result = append(result,nums[i+m])
		}
	}
	return result
}
// @lc code=end

