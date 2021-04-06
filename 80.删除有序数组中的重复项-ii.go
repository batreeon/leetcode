/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除排序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	n := len(nums)
	for i := 0 ; i < n ; {
		j := i+1
		for ; j < n ; j++ {
			if nums[j] != nums[i] {
				if j - 2 > i {
					copy(nums[i+2:],nums[j:])
					n = n - (j-(i+2))
					i = i + 2
					break
				}
				i = j
				break
			}
		}
		// 最后一串只有相同的数字
		// 若长度大于2，那么就更新长度
		// 由于j越界了，那么就是所有的数字都访问了，那么i也应该退出
		if j == n {
			if j-i >= 2 {
				n = i+2
			}
			i = n
		}
	}
	return n
}
// @lc code=end

