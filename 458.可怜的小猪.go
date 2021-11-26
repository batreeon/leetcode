/*
 * @lc app=leetcode.cn id=458 lang=golang
 *
 * [458] 可怜的小猪
 */

// @lc code=start
package main

import "math"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	jinzhi := minutesToTest / minutesToDie + 1
	var result int
	for ; math.Pow(float64(jinzhi), float64(result)) < float64(buckets); result++ {}
	return result
}
// @lc code=end

