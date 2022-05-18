/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	for ; n >= 10; {
		n = ishappy(n)
	}
	if n == 1 || n == 7 {
		return true
	}
	return false
}
func ishappy(n int) int {
	if n < 10 {
		return n
	}
	result := 0
	for ; n > 0; n /= 10 {
		num := (n % 10)
		result += num * num
	}
	return result
}
// @lc code=end

