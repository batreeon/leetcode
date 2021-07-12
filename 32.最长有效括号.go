/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
func longestValidParentheses(s string) int {
	/*
	if len(s) == 0 {
		return 0
	}
	stack := []int{}
	stack = append(stack,-1)
	result := 0
	for i := range s {
		if s[i] == '(' {
			stack = append(stack,i)
		}else{
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack,i)
			}else{
				if i - stack[len(stack)-1] > result {
					result = i - stack[len(stack)-1]
				}
			}
		}
	}
	return result
	*/
	//动规
	dp := make([]int,len(s))
	result := 0
	// dp[i]表示以s[i]结尾的最长有小括号，因此若s[i]=='(',则dp[i]一定为0
	for i := 1 ; i < len(s) ; i++ {
		if s[i] == ')' {
			// ()
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				}else{
					dp[i] = 2
				}
			// ))
			}else if i - dp[i-1] > 0 && s[i - dp[i-1] - 1] == '(' {
				if i-dp[i-1]-1 >= 2 {
					// ()((..))
					dp[i] = dp[i-1]+2 + dp[i-dp[i-1]-2]
				}else{
					dp[i] = dp[i-1]+2
				}
			}
			if dp[i] > result {
				result = dp[i]
			}
		}
	}
	return result
}
// @lc code=end

