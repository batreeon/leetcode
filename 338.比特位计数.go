/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(num int) []int {
	ans := []int{0}
	if num == 0 {
		return ans
	}
	back := 1
	for i,j := 1,back ; i <= num ; i++ {
		ans = append(ans,ans[i-back]+1)
		j--
		if j == 0 {
			back *= 2
			j = back
		}
	}
	return ans
}
// @lc code=end

