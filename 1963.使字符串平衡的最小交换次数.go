/*
 * @lc app=leetcode.cn id=1963 lang=golang
 *
 * [1963] 使字符串平衡的最小交换次数
 */

// @lc code=start
package main

func minSwaps(s string) int {
	result, left, right := 0, 0, 0
	for _, b := range s {
		if b == '[' {
			left++
		} else {
			right++
		}
		if right > left {
			// 右括号多余左括号，交换一次，每次交换会造成left增一，right减一
			result++
			left++
			right--
		}
	}
	return result
}

// @lc code=end
