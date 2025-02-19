/*
 * @lc app=leetcode.cn id=2595 lang=golang
 *
 * [2595] 奇偶位数
 */

// @lc code=start
func evenOddBit(n int) []int {
    res := []int{0,0}
	i := 0
	for n != 0 {
		if n & 1 == 1{
			res[i%2]++
		}
		n >>= 1
		i++
		i %= 2
	}
	return res
}
// @lc code=end

