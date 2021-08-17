/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] 等差数列划分
 */

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	result := 0
	for l := 0 ; l <= n-3 ; {
		// l <= n-3，长度小于3后面末尾就不考虑了
		r := l + 2
		// r从l+2开始考虑，长度至少为3
		for ; r < n ; r++ {
			// 不符合等差数列了，直接退出
			if nums[r]-nums[r-1] != nums[r-1]-nums[r-2] {
				break
			}
		}
		// r-1 - l + 1 有效的长度
		length := r-l
		if length >= 3 {
			// 子数组，分别考察长度为3、4、5...的子数组的个数
			for i := 2 ; length - i > 0 ; i++ {
				result += length - i
			}
		}
		// 更新左边界
		l = r - 1
	}
	return result
}
// @lc code=end

