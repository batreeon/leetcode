/*
 * @lc app=leetcode.cn id=1824 lang=golang
 *
 * [1824] 最少侧跳次数
 */

// @lc code=start
func minSideJumps(obstacles []int) int {
	n := len(obstacles)
	if n == 2 {
		return 0
	}

	bp := make([]int, 4)
	bp[1], bp[3] = 1, 1

	for i := 1; i < n-1; i++ {
		for j := 1; j < 4; j++ {
			if j == obstacles[i] {
				bp[j] = -1
				continue
			}
			if bp[j] == -1 {
				bp[j] = -2
				continue
			}
		}
		for j := 1; j < 4; j++ {
			if bp[j] == -2 {
				// 已经考虑了obstacles[i]为0的情况
				bp[j] = mins(bp[1], bp[2], bp[3]) + 1
			}
		}
	}
	return mins(bp[1], bp[2], bp[3])
}

func mins(x, y, z int) int {
	return min(min(x, y), min(y, z))
}

func min(x, y int) int {
	if x < 0 {
		return y
	}
	if y < 0 {
		return x
	}
	if x < y {
		return x
	}
	return y
}
// @lc code=end

