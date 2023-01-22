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

	starts := make([]int, 4)
	starts[1], starts[3] = 1, 1
	ends := make([]int, 4)

	for i := 1; i < n-1; i++ {
		for j := 1; j < 4; j++ {
			if j == obstacles[i] {
				ends[j] = -1
				continue
			}
			if starts[j] == -1 {
				ends[j] = -2
				continue
			}
			ends[j] = starts[j]
		}
		for j := 1; j < 4; j++ {
			if ends[j] == -2 {
				// 已经考虑了obstacles[i]为0的情况
				ends[j] = mins(ends[1], ends[2], ends[3]) + 1
			}
		}
		copy(starts, ends)
	}
	return mins(starts[1], starts[2], starts[3])
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

