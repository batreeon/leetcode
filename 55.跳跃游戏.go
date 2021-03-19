/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	length := len(nums)
	dp := make([]bool,length)
	dp[length-1] = true
	for i := length-2 ; i >= 0 ; i-- {
		for j := nums[i] ; j > 0 ; j-- {
			if i + j >= length - 1{
				dp[i] = true
				break
			}else{
				if dp[i+j] == true {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[0]
}
// @lc code=end

