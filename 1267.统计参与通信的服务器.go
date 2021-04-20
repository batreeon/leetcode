/*
 * @lc app=leetcode.cn id=1267 lang=golang
 *
 * [1267] 统计参与通信的服务器
 */

// @lc code=start
func countServers(grid [][]int) int {
	// 从前到后扫描一遍，再从后到前扫描一遍
	// 扫描一遍可能出现该行及该列第一个出现的元素没办法及时找到连接对象
	// 因此扫描从不同方向扫描两遍，若两个方向都没找到连接对象，那么这个点就无法进行连接
	// 用总服务器数减去无连接的服务器数就是答案
	r,c := len(grid),len(grid[0])
	unlinked1,unlinked2 := make(map[int]bool),make(map[int]bool)
	linkedrows1,linkedcols1 := make(map[int]bool),make(map[int]bool)
	linkedrows2,linkedcols2 := make(map[int]bool),make(map[int]bool)
	totalServer := 0
	for i := 0 ; i < r ; i++ {
		for j := 0 ; j < c ; j++ {
			if grid[i][j] == 1 {
				totalServer++
				if !linkedrows1[i] && !linkedcols1[j] {
					unlinked1[i*c+j] = true
				}
				linkedrows1[i] = true
				linkedcols1[j] = true
			}
		}
	}
	for i := r-1 ; i >= 0 ; i-- {
		for j := c-1 ; j >= 0 ; j-- {
			if grid[i][j] == 1 {
				if !linkedrows2[i] && !linkedcols2[j] {
					unlinked2[i*c+j] = true
				}
				linkedrows2[i] = true
				linkedcols2[j] = true
			}
		}
	}

	unlinkedNum := 0
	for k1 := range unlinked1 {
		if unlinked2[k1] {
			unlinkedNum++
		}
	}
	return totalServer - unlinkedNum
}
// @lc code=end

