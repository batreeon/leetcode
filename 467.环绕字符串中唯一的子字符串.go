/*
 * @lc app=leetcode.cn id=467 lang=golang
 *
 * [467] 环绕字符串中唯一的子字符串
 */

// @lc code=start
// 动态规划
// 建立数组dp[26]，数组元素dp[c]表示以字符c结尾的子字符串的个数
func findSubstringInWraproundString(p string) int {
	dp := make([]int, 26)
	dp[p[0]-'a']++
	// j用来记录当前连续的子串长度
	// 比如dabc，当遍历到a，j为1；遍历到b，j为2；遍历到c，j为3
	for i, j := 1, 1; i < len(p); i++ {
		c := p[i] - 'a';
		preC := p[i-1] - 'a'
		if (preC == 25 && c == 0) || (c == preC + 1) {
			j++
		}else{
			j = 1
		}
		dp[c] = max(dp[c], j)
	}
	result := 0
	for _, num := range dp {
		result += num
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

