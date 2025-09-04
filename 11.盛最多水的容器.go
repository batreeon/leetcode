/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
    l, r := 0, len(height)-1
	result := -1

	for l < r {
		result = max(result, min(height[l], height[r]) * (r-l))
		if height[l] < height[r] {
			l++
		}else{
			r--
		}
	}

	return result
}
// @lc code=end

