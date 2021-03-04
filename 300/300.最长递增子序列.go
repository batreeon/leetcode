/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
package main
func lengthOfLIS(nums []int) int {
	//分治
	l := len(nums)
	f := make([]int,l)
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	ans := 0
	for i := 0 ; i < l ; i++ {
		maxf := 0
		for j := 0 ; j < i ; j++ {
			if nums[j] < nums[i] {
				maxf = max(maxf,f[j])
			}
		}
		f[i] = maxf + 1
		ans = max(ans,f[i])
	}
	return ans
}
// @lc code=end

