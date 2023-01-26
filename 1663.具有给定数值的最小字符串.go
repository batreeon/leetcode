/*
 * @lc app=leetcode.cn id=1663 lang=golang
 *
 * [1663] 具有给定数值的最小字符串
 */

// @lc code=start
func getSmallestString(n int, k int) string {
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = 'a'
	}
	k -= n 
	for i := n-1; i >= 0; i-- {
		if k == 0 {
			break
		}
		volume := int('z' - result[i])
		if volume > k {
			result[i] += byte(k)
			k -= k
		}else{
			result[i] = 'z'
			k -= volume
		}
	}
	return string(result)
}
// @lc code=end

