/*
 * @lc app=leetcode.cn id=740 lang=golang
 *
 * [740] 删除并获得点数
 */

// @lc code=start
func deleteAndEarn(nums []int) int {
	cnt := map[int]int{}
	dp := []int{0}
	sort.Ints(nums)

	for i := 0 ; i < len(nums) ; i++ {
		cnt[nums[i]]++
		if nums[i] != dp[len(dp)-1] {
			dp = append(dp,nums[i])
		}
	}

	last := dp[1]
	dp[1] = dp[1]*cnt[dp[1]]
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 2 ; i < len(dp) ; i++ {
		if dp[i] - last == 1 {
			last = dp[i]
			dp[i] = max(dp[i-1],dp[i-2]+dp[i]*cnt[dp[i]])
		}else{
			last = dp[i]
			dp[i] = dp[i-1]+dp[i]*cnt[dp[i]]
		}
	}
	return dp[len(dp)-1]
}
// @lc code=end

