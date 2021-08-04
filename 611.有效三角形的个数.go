/*
 * @lc app=leetcode.cn id=611 lang=golang
 *
 * [611] 有效三角形的个数
 */

// @lc code=start
package main
import "sort"

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	result := 0
	for idx1 := 0 ; idx1 < n-2 ; idx1++ {
		for idx2 := idx1+1 ; idx2 < n-1 ; idx2++ {
			// sort.Search返回的是在nums[idx2+1:]子数组中的下标
			// 加上idx2+1,就得到了在nums中的下标
			// 两小边之和大于第三边
			idx3 := sort.Search(n-1-idx2,func(i int) bool {
				return nums[idx1] + nums[idx2] <= nums[idx2+1:][i]
			}) + 1 + idx2
			// 前两条边长分别是nums[idx1],nums[idx2],
			// idx2+1到idx3-1之间这些都可以作为第三条边，加上就行了
			result += idx3-idx2-1
		}
	}
	return result
}
// @lc code=end

