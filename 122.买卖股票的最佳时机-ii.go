/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	var ans int
	for i := 0;i < len(prices);i++ {
		buy := prices[i]
		for ;i < len(prices)-1 && prices[i+1] > prices[i]; {
			i++
		}
		ans += prices[i]-buy
	}
	return ans
}
// @lc code=end

