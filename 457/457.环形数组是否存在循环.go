/*
 * @lc app=leetcode.cn id=457 lang=golang
 *
 * [457] 环形数组是否存在循环
 */

// @lc code=start
package main

func circularArrayLoop(nums []int) bool {
	n := len(nums)
	seen := map[int]bool{}
	for i := 0 ; i < n ; i++ {
		if seen[i] {
			// 已经考察过了，不能构成循环
			continue
		}
		seen[i] = true
		// 记录符号，同一轮移动方向要一致
		sign := 1
		if nums[i] < 0 {
			sign = -1
		}
		// 当前这一路径
		circle := map[int]bool{i:true}
		nextIdx := i
		for {
			lastIdx := nextIdx
			nextIdx = (nextIdx + nums[nextIdx])%n
			if nextIdx < 0 {
				// 计算出的下标为负，找到相应的下标
				nextIdx += n
			}
			if nums[nextIdx] * sign < 0 {
				// 同一条路径上出现了方向不同的，不可行
				break
			}
			if lastIdx == nextIdx {
				// 原地跳，当然不能构成循环了
				break
			}
			if circle[nextIdx] {
				// 由上一个下标跳到的下一个下标，在当前路径上，则可以构成循环
				return true
			}
			// 记录一下，出现在了当前路径上，已经考察过了
			circle[nextIdx] = true
			seen[nextIdx] = true
		}
	}
	return false
}
// @lc code=end

