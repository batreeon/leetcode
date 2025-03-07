/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	// 从0开始做起点，依次尝试
	for i := 0; i < len(gas); {
		// 当i为起点时，如果不能到达next下标，那么下一次应该直接从next下标开始尝试，
		// 而不是从i+1开始尝试
		next, ok := circle(gas, cost, i)
		if ok {
			return i
		}
		if next > i {
			i = next
		}else{
			return -1
		}
	}
	return -1
}

func circle(gas []int, cost []int, start int) (int, bool) {
	g := gas[start]
	for i := (start+1)%len(gas);; i = (i+1)%len(gas) {
		last := i-1
		if i == 0 {
			last = len(gas)-1
		}
		if g - cost[last] < 0 {
			return i, false
		}
		if i == start {
			return 0, true
		}
		g = g - cost[last] + gas[i]
	}
}
// @lc code=end

