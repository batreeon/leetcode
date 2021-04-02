/*
 * @lc app=leetcode.cn id=1306 lang=golang
 *
 * [1306] 跳跃游戏 III
 */

// @lc code=start
func canReach(arr []int, start int) bool {
	// 若访问过就记录一下，不然可能会死循环？
	seen := map[int]bool{}
	n := len(arr)
	// 将不可行的记下来，重用
	can := map[int]bool{}
	var canreach func(start int) bool
	canreach = func(start int) bool {
		seen[start] = true
		if v, ok := can[start]; ok {
			return v
		}
		left, right := start-arr[start], start+arr[start]
		if left < 0 && right >= n {
			can[start] = false
			return false
		}

		// 这两个检查可以合并到后面两个if中，但这两个成本貌似低一点，
		// 但好像没啥差别，离最终结果的都很近
		if left >= 0 && arr[left] == 0 {
			return true
		}
		if right < n && arr[right] == 0 {
			return true
		}

		if left >= 0 {
			if seen[left] == false {
				if canreach(left) {
					return true
				}
			}
		}
		if right < n {
			if seen[right] == false {
				if canreach(right) {
					return true
				}
			}
		}
		return false
	}

	return canreach(start)
}

// @lc code=end

