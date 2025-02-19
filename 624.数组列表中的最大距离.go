/*
 * @lc app=leetcode.cn id=624 lang=golang
 *
 * [624] 数组列表中的最大距离
 */

// @lc code=start
func maxDistance(arrays [][]int) int {
    mini, maxi := 0, 0
	minv, maxv := arrays[0][0], arrays[0][len(arrays[0])-1]
	for i := 1; i < len(arrays); i++ {
		if arrays[i][0] < arrays[mini][0] {
			mini = i
			minv = arrays[i][0]
		}
		if arrays[i][len(arrays[i]) - 1] > arrays[maxi][len(arrays[maxi]) - 1] {
			maxi = i
			maxv = arrays[i][len(arrays[i]) - 1]
		}
	}

	res := -2 * 10000 - 1
	for i := 0; i < len(arrays); i++ {
		if i != mini {
			res = max(res, abs(arrays[i][len(arrays[i]) - 1] - minv))
		}
		if i != maxi {
			res = max(res, abs(maxv - arrays[i][0]))
		}
	}
	return res
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

