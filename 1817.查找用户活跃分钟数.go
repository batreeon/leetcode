/*
 * @lc app=leetcode.cn id=1817 lang=golang
 *
 * [1817] 查找用户活跃分钟数
 */

// @lc code=start
func findingUsersActiveMinutes(logs [][]int, k int) []int {
	// 记录每个id的活跃分钟
	activeMinutes := make(map[int]map[int]struct{})
	for _, log := range logs {
		id := log[0]
		if _, ok :=  activeMinutes[id]; !ok {
			activeMinutes[id] = make(map[int]struct{})
		}
		activeMinutes[id][log[1]] = struct{}{}
	}

	answer := make([]int, k)
	for _, activeMinute := range activeMinutes {
		answer[len(activeMinute)-1]++
	}

	return answer
}
// @lc code=end

