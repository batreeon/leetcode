/*
 * @lc app=leetcode.cn id=526 lang=golang
 *
 * [526] 优美的排列
 */

// @lc code=start
func countArrangement(n int) int {
	result := 0
	// available[i]存第i为可以放的数
	available := make([][]int,n+1)
	for i := 1 ; i <= n ; i++ {
		available[i] = []int{}
		for j := 1 ; j <= n ; j++ {
			if j%i == 0 || i%j == 0 {
				available[i] = append(available[i],j)
			}
		}
	}
	seen := make([]bool,n+1)
	var tb func(idx int)
	tb = func(idx int) {
		if idx == n + 1 {
			result++
			return
		}
		for _,j := range available[idx] {
			if !seen[j] {
				seen[j] = true
				tb(idx+1)
				seen[j] = false
			}
		}
	}

	tb(1)
	return result
}
// @lc code=end

