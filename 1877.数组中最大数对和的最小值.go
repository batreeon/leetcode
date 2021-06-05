/*
 * @lc app=leetcode.cn id=1877 lang=golang
 *
 * [1877] 数组中最大数对和的最小值
 */

// @lc code=start
package main
func minPairSum(nums []int) int {
	sort.Ints(nums)
	result := 0
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i,j := 0,len(nums)-1 ; i < j ; i,j = i+1,j-1 {
		result = max(result,nums[i]+nums[j])
	}
	return result
}
// @lc code=end

