/*
 * @lc app=leetcode.cn id=1893 lang=golang
 *
 * [1893] 检查是否区域内所有整数都被覆盖
 */

// @lc code=start
func isCovered(ranges [][]int, left int, right int) bool {
	m := make(map[int]bool)
	for i := left ; i <= right ; i++ {
		m[i] = true
	}
	total := right-left+1
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	for _,rge := range ranges {
		if rge[0] > right || rge[1] < left {
			continue
		}
		l,r := max(left,rge[0]),min(right,rge[1])
		for i := l ; i <= r ; i++ {
			if m[i] {
				total--
				m[i] = false
			}
		}
	}
	return total == 0
}
// @lc code=end

