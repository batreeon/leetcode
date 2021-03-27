/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] 交错字符串
 */

// @lc code=start
package main
func isInterleave(s1 string, s2 string, s3 string) bool {
	/*
	// dfs，果然是超时的，要用动规
	var isinterleave func(first bool , s1,s2,s3 string) bool 
	isinterleave = func(first bool , s1,s2,s3 string) bool {
		if len(s3) == 0 {
			return true
		}
		var ans,ans1,ans2 bool
		if first {
			for i := 0 ; i < len(s1) ; i++ {
				if s3[:i+1] == s1[:i+1] {
					ans1 = isinterleave(false,s2,s1[i+1:],s3[i+1:])
					if ans1 {
						return true
					}
				}else{
					break
				}
			}
			for i := 0 ; i < len(s2) ; i++ {
				if s3[:i+1] == s2[:i+1] {
					ans2 = isinterleave(false,s1,s2[i+1:],s3[i+1:])
					if ans2 {
						return true
					}
				}else{
					break
				}
			}
			return ans1 || ans2
		}else{
			for i := 0 ; i < len(s1) ; i++ {
				if s3[:i+1] == s1[:i+1] {
					ans = isinterleave(false,s2,s1[i+1:],s3[i+1:])
					if ans {
						return true
					}
				}else{
					break
				}
			}
		}
		return false
	}

	if len(s1) + len(s2) != len(s3) {
		return false
	}
	return isinterleave(true,s1,s2,s3)
	*/

	// 来点动规吧
	if len(s1) + len(s2) != len(s3) {
		return false
	}
	r,c := len(s1),len(s2)
	f := make([][]bool,r+1)
	// 第一行表示s2的前j个字符和s3的前j个字符是否匹配
	// 第一列表示s1的前i个字符和s3的前i个字符是否匹配
	// f[1][1]看f[1][0]和s2[0],与f[0][1]和s1[0]
	// f[1][2]看f[1][1]和s2[1],与f[0][2]和s1[0]
	// 。。。
	// 不用考虑什么交替性，始终是交替的，只要考察是否能组成就可以了，
	// 所谓的非交替，那么连续来自同一字符串的，不就可以合并为一组提交吗，真是笨
	
	for i := 0 ; i <= r ; i++ {
		f[i] = make([]bool,c+1)
		if i == 0 {
			f[0][0] = true
		}
		for j := 0 ; j <= c ; j++ {
			if i > 0 {
				f[i][j] = f[i][j] || (f[i-1][j] && s1[i-1] == s3[i+j-1])
			}
			if j > 0 {
				f[i][j] = f[i][j] || (f[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}
	return f[r][c]
}
// @lc code=end

