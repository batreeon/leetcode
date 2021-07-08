/*
 * @lc app=leetcode.cn id=930 lang=golang
 *
 * [930] 和相同的二元子数组
 */

// @lc code=start
package main
func numSubarraysWithSum(nums []int, goal int) int {
	sum := 0
	m := map[int]int{}
	m[0] = 1
	for i := 0 ; i < len(nums) ; i++ {
		sum += nums[i]
		if _,ok := m[sum] ; ok {
			m[sum]++
		}else{
			m[sum] = 1
		}
	}
	result := 0
	for k,v := range m {
		if k < goal {
			continue
		}
		target := k-goal
		if vtarget,ok := m[target] ; ok {
			if k == target {
				// 这种情况就是goal是零，那么就是有连续v-1个0出现
				// 其子数组和为0，就是(v-1)*(v-1+1)/2
				result += v*(v-1)/2
			}else{
				result += v*vtarget
			}
		}
	}
	return result
}
// @lc code=end

