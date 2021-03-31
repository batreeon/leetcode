/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	/*
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	n := len(nums)
	f := make([]int,n)
	f[0] = nums[0]
	result := f[0]
	for i := 1 ; i < n ; i++ {
		f[i] = max(nums[i],f[i-1]*nums[i])
		result = max(result,f[i])
	}
	return result
	*/
	
	// 上面这样简单地考虑由前一位转换而来是不够的，
	// 若当前位为负数，那么就是期望前面的最小负乘积了
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	fmin,fmax := nums[0],nums[0]
	n := len(nums)
	result := nums[0]
	for i := 1 ; i < n ; i++ {
		// 用下面这种解法，就不需要用临时变量存储上一步的fmin和fmax了
		fmin,fmax = min(min(fmin*nums[i],fmax*nums[i]),nums[i]),max(max(fmax*nums[i],fmin*nums[i]),nums[i])
		result = max(result,fmax)
	}
	return result
}
// @lc code=end