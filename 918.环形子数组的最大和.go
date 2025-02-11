/*
 * @lc app=leetcode.cn id=918 lang=golang
 *
 * [918] 环形子数组的最大和
 */

// @lc code=start
func maxSubarraySumCircular(nums []int) int {
    if len(nums) == 1 {
		return nums[0]
	}
    maxSum := nums[0]
	maxPre := maxSum
	minSum := nums[0]
	minPre := minSum

	sum := nums[0]

	for _, num := range nums[1:] {
		if maxPre < 0 {
			maxPre = num
		}else{
			maxPre += num
		}
		if maxPre > maxSum {
			maxSum = maxPre
		}

		if minPre > 0 {
			minPre = num
		}else{
			minPre += num
		}
		if minPre < minSum {
			minSum = minPre
		}
		sum += num
	}
	if maxSum < 0 {
		return maxSum
	}
	if maxSum > sum - minSum {
		return maxSum
	}
	return sum - minSum
}
// @lc code=end

