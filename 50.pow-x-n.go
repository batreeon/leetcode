/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	// 快速幂
	is_minus := n < 0
	power := 1.0
	if n < 0 {
		n = -1 * n
	}
	for ; n > 0 ; n >>= 1 {
		if n&1 == 1 {
			// 若指数为奇数
			power *= x
		}
		x *= x
	}
	if is_minus {
		power = 1/power
	}
	return power
}
// @lc code=end

