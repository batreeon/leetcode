/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(n int) int {
	// if n == 0 {
	// 	return 0
	// }
	// if n == 1 {
	// 	return 1
	// }
	if n < 2 {
		return n
	}
	f0,f1 := 0,1
	for i := 2 ; i <= n ; i++ {
		f0,f1 = f1,f0+f1
	}
	return f1
}
// @lc code=end

