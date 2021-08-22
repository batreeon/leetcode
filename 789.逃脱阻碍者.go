/*
 * @lc app=leetcode.cn id=789 lang=golang
 *
 * [789] 逃脱阻碍者
 */

// @lc code=start
func escapeGhosts(ghosts [][]int, target []int) bool {
	// 计算幽灵到target的距离（步数），及从起点到target的距离，
	// 如果到target时幽灵都抓不到，那么前面一定也抓不到
	// 这个具体的证明过程我有点不太会
	pathLen := distance([]int{0,0},target)
	for _,ghost := range ghosts {
		if distance(ghost,target) <= pathLen {
			return false
		}
	}
	return true
}
func distance(p1,p2 []int) int {
	return abs(p1[0]-p2[0])+abs(p1[1]-p2[1])
} 
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
// @lc code=end

