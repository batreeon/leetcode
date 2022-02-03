/*
 * @lc app=leetcode.cn id=1414 lang=golang
 *
 * [1414] 和为 K 的最少斐波那契数字数目
 */

// @lc code=start
func findMinFibonacciNumbers(k int) int {
	fib := []int{1,1}
	for fib[len(fib)-1] < k {
		fib = append(fib, fib[len(fib)-1] + fib[len(fib)-2])
	}
	result := 0
	for i := len(fib)-1; k > 0; i-- {
		if k >= fib[i] {
			k -= fib[i]
			result++
		}
	}
	return result
}
// @lc code=end

