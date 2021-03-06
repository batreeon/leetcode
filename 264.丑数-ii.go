/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	nums := []int{1}
	i2,i3,i5 := 0,0,0
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	for i := 1 ; i < n ; i++ {
		v := min( min(2*nums[i2],3*nums[i3]) , 5*nums[i5] )
		nums = append(nums,v)
		for 2 * nums[i2] <= v {
			i2++
		}
		for 3 * nums[i3] <= v {
			i3++
		}
		for 5 * nums[i5] <= v {
			i5++
		}
	}
	return nums[n-1]
}
// @lc code=end

