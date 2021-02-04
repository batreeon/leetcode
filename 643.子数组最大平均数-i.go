/*
 * @lc app=leetcode.cn id=643 lang=golang
 *
 * [643] 子数组最大平均数 I
 */

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	if len(nums) == 0 || len(nums) < k {
		return 0
	}
	p,max := 0.0,0.0
	i := 0
	kf := float64(k)
	for ; i < k ; i++ {
		p += float64(nums[i])/kf
	}
	max = p 
	for ; i < len(nums) ; i++ {
		p -= float64(nums[i-k])/kf
		p += float64(nums[i])/kf
		if p > max {
			max = p
		}
	}
	return max
}
// @lc code=end

