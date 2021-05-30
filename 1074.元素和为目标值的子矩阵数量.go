/*
 * @lc app=leetcode.cn id=1074 lang=golang
 *
 * [1074] 元素和为目标值的子矩阵数量
 */

// @lc code=start
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	var r,c int
	r = len(matrix)
	if r != 0 {
		c = len(matrix[0])
	}
	// 前缀和
	sum := make([][]int,r+1)
	sum[0] = make([]int,c+1)
	for	i := 1 ; i <= r ; i++ {
		sum[i] = make([]int,c+1)
		for j := 1 ; j <= c ; j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	result := 0
	// 上界
	for top := 1 ; top <= r ; top++ {
		// 下界
		for bot := top ; bot <= r ; bot++ {
			cur := 0
			// 哈希记录前缀和
			m := make(map[int]int)
			for i := 1 ; i <= c ; i++ {
				cur = sum[bot][i] - sum[top-1][i]
				// 该子矩阵和等于target
				if cur == target {
					result++
				}
				// 前面有子矩阵的和 = cur-target
				if v,ok := m[cur-target] ; ok {
					result += v
				}
				m[cur]++
			}
		}
	}
	return result
}
// @lc code=end

