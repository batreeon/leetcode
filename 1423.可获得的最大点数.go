/*
 * @lc app=leetcode.cn id=1423 lang=golang
 *
 * [1423] 可获得的最大点数
 */

// @lc code=start
func maxScore(cardPoints []int, k int) int {
	right := k-1
	maxSum := 0
	for i := 0 ; i <= right ; i++ {
		maxSum += cardPoints[i]
	}
	nowSum := maxSum
	right--
	left := len(cardPoints)-1
	for ; right >= -1 ; {//right=-1时，就是只取末尾
		nowSum -= cardPoints[right+1]
		nowSum += cardPoints[left]
		if maxSum < nowSum {
			maxSum = nowSum
		}
		right--
		left--
	}
	return maxSum
}
// @lc code=end

