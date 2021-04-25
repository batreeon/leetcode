/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {

	/*
	n := len(s)
	dp := make([]int,n+1)
	dp[n] = 1
	for i := n-1 ; i >= 0 ; i-- {
		if s[i] == '0' {
			continue
		}
		// s[i] != '0',那么单独的s[i]一定可以被解码，则至少有dp[i+1]种解码
		res := dp[i+1]
		for j := i+1 ; j < n ; j++ {
			// 当j==i+1时，s[i:j]就是s[i],这个前面已经考虑过了
			num,_ := strconv.Atoi(s[i:j+1])
			if num > 0 && num <= 26 {
				res += dp[j+1]
			}
			// 再多一位肯定比26大了，没必要看了
			if num >= 3 {
				break
			}
		}
		dp[i] = res
	}

	return dp[0]
	*/

	n := len(s)
	dp1,dp2 := 1,1
	for i := n-1 ; i >= 0 ; i-- {
		if s[i] == '0' {
			dp1,dp2 = 0,dp1
			continue
		}
		dp := dp1
		if i != n-1 {
			num,_ := strconv.Atoi(s[i:i+2])
			if num <= 26 {
				dp += dp2
			}
		}
		dp1,dp2 = dp,dp1
	}

	return dp1
}
// @lc code=end

