/*
 * @lc app=leetcode.cn id=377 lang=golang
 *
 * [377] 组合总和 Ⅳ
 */

// @lc code=start
func combinationSum4(nums []int, target int) int {
	// 为什么这样解把顺序也考虑到了呢
	dp := make([]int,target+1)
	dp[0] = 1
	for i := 1 ; i < target+1 ; i++ {
		for _,num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
// @lc code=end

