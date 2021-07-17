func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1 ; i < len(nums) ; i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}//如果小了，那么就没必要加了，他对找最大值没有贡献，反倒拖了后腿
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}