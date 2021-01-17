/*
 * @lc app=leetcode.cn id=402 lang=golang
 *
 * [402] 移掉K位数字
 */

// @lc code=start
func removeKdigits(num string, k int) string {
	if k == len(num) {
		return string('0')
	}
	for i := 0 ; i < k ; i++ {
		for j := 0 ; j < len(num); j++ {
			if j == len(num) - 1 {
				num = num[:len(num)-1]
				break
			}
			if num[j+1] < num[j] {
				num = num[:j] + num[j+1:]
				break
			}
		}
	}
	for ;num[0] == '0' && len(num) > 1; {
		num = num[1:]
	}
	return num
}
// @lc code=end

