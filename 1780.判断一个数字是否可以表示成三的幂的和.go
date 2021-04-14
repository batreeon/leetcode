/*
 * @lc app=leetcode.cn id=1780 lang=golang
 *
 * [1780] 判断一个数字是否可以表示成三的幂的和
 */

// @lc code=start
func checkPowersOfThree(n int) bool {
	for {
		if n == 1 {
			return true
		}
		if n % 3 == 0 {
			n /= 3
		}else if (n-1) % 3 == 0 {
			n = (n-1) / 3
		}else{
			return false
		}
	}
}
// @lc code=end

