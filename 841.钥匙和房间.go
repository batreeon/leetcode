/*
 * @lc app=leetcode.cn id=841 lang=golang
 *
 * [841] 钥匙和房间
 */

// @lc code=start
func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)

	seen := map[int]bool{}
	var dfs func(v int)
	dfs = func(v int) {
		seen[v] = true
		for _,next := range rooms[v] {
			if !seen[next] {
				dfs(next)
			}
		}
	}

	dfs(0)

	return len(seen) == n
}
// @lc code=end

