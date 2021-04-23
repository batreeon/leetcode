/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	nums := map[int]bool{}
	for i := 1 ; i <= 10 ; i++ {
		nums[i*i] = true
	}

	var numsquares func(n int) int 
	numsquares = func(n int) int {
		if nums[n] {
			return 1
		}
		square := 1
		for (square+1)*(square+1) <= n {
			square++
		}

		result := 0
		for i := 
	}
}
// @lc code=end

