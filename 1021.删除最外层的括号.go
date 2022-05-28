/*
 * @lc app=leetcode.cn id=1021 lang=golang
 *
 * [1021] 删除最外层的括号
 */

// @lc code=start
func removeOuterParentheses(s string) string {
	status := 0
	result := []byte{}
	for _, b := range s {
		if b == '(' {
			if status != 0 {
				result = append(result, '(')
			}
			status--
		}else{
			if status != -1 {
				result = append(result, ')')
			}
			status++
		}
	}
	return string(result)
}
// @lc code=end

