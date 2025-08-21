/*
 * @lc app=leetcode.cn id=1504 lang=golang
 *
 * [1504] 统计全 1 子矩形
 */

// @lc code=start
func numSubmat(mat [][]int) int {
	r, c := len(mat), len(mat[0])
	result := 0
	for top := 0; top < r; top++ {
		sub := make([]int, c)
		for bottom := top; bottom < r; bottom++ {
			h := bottom - top + 1
			last := -1
			for i := 0; i < c; i++ {
				sub[i] += mat[bottom][i]
				if sub[i] != h {
					last = i
				} else {
					result += i - last
				}
			}
		}
	}
	return result
}

// @lc code=end

