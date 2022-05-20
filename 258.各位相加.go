/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */

// @lc code=start
func addDigits(num int) int {
	for ; num >= 10; {
		add := 0
		for ; num > 0; num /= 10 {
			add += num % 10
		}
		num = add
	}
	return num
}
// @lc code=end

