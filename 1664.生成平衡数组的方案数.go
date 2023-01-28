/*
 * @lc app=leetcode.cn id=1664 lang=golang
 *
 * [1664] 生成平衡数组的方案数
 */

// @lc code=start
func waysToMakeFair(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 1
	}
	if l == 2 {
		return 0
	}
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	sums[1] = nums[1]
	for i := 2; i < l; i++ {
		sums[i] = sums[i-2] + nums[i]
	}
	var totaleven, totalodd int 
	if (l-1) % 2 == 0 {
		totaleven = sums[l-1]
		totalodd = sums[l-2]
	}else{
		totaleven = sums[l-2]
		totalodd = sums[l-1]
	}

	var evensum, oddsum int
	result := 0
	for i := 0; i < l; i++ {
		if i == 0 {
			evensum = totalodd
			oddsum = totaleven - nums[0]
		}else if i == 1 {
			evensum = nums[0] + totalodd - nums[1]
			oddsum = totaleven - nums[0]
		}else{
			if i % 2 == 0 {
				evensum = sums[i-2] + totalodd - sums[i-1]
				oddsum = sums[i-1] + totaleven - sums[i]
			}else{
				evensum = sums[i-1] + totalodd - sums[i]
				oddsum = sums[i-2] + totaleven - sums[i-1]
			}
		}
		if evensum == oddsum {
			result++
		}
	}

	return result
}
// @lc code=end

