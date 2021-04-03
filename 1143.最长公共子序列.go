/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {

	// 若text1[i] == text2[j],那么f[i][j] = f[i-1][j-1]+1
	// 若不等，则应该取f[i-1][j]和f[i][j-1]中的最大值
	l1,l2 := len(text1),len(text2)
	f := make([]int,l2+1)
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	result := 0
	for i := 0 ; i < l1 ; i++ {
		curf := make([]int,l2+1)
		for j := 1 ; j <= l2 ; j++ {
			if text1[i] == text2[j-1] {
				curf[j] = f[j-1] + 1
			}else{
				curf[j] = max(f[j],curf[j-1])
			}
			result = max(result,curf[j])
		}
		copy(f,curf)
	}
	return result
	// 为什么不能用一个f：
	// abcba
	// abcbcba
}
// @lc code=end

