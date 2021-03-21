/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
package main
import "sort"

func fourSum(nums []int, target int) [][]int {
	// 唉，和三数之和用的同样的方法
	l := len(nums)
	sort.Ints(nums)
	ans := [][]int{}
	for first := 0 ; first < l-3 ; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		res1 := target - nums[first]
		for second := first+1 ; second < l-2 ; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			res2 := res1 - nums[second]
			for third := second + 1 ; third < l-1 ; third++ {
				if third > second+1 && nums[third] == nums[third-1] {
					continue
				}
				res3 := res2 - nums[third]
				fifth := l-1
				for fifth > third && nums[fifth] > res3 {
					fifth--
				}
				if fifth == third {
					break
				}
				if nums[fifth] == res3 {
					ans = append(ans,[]int{nums[first],nums[second],nums[third],nums[fifth]})
				}
			}
		}
	}
	return ans
}
// @lc code=end

