/*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 */

// @lc code=start
package main
import "sort"
func findUnsortedSubarray(nums []int) int {
	numsCopy := make([]int,len(nums))
	copy(numsCopy,nums)
	sort.Ints(numsCopy)
	left,right := 0,len(nums)-1
	for ; left <= right ; left++ {
		if numsCopy[left] != nums[left] {
			break
		}
	}
	for ; right >= left ; right-- {
		if numsCopy[right] != nums[right] {
			break
		}
	}
	return right-left+1 
}
// @lc code=end

