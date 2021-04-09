/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除排序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	l := len(nums)
	for i := 0 ; i < l ; i++ {
		// 先判断是否有重复元素
		if i != 0 && nums[i] == nums[i-1] {
			// i指示第二个重复元素
			j := i+1
			for  ; j < l && nums[i] == nums[j] ; j++ {}
			// j指示下一个新元素
			// 后面的元素全部前移，覆盖掉多余的元素
			copy(nums[i+1:],nums[j:l])
			l -= j-i-1
			// i始终维护着已处理好,无多余元素的序列的末尾
		}
	}
	// 有一个问题，后面的元素，会因为前面的元素迁移而跟着多次迁移，有点浪费
	return l
}
// @lc code=end

