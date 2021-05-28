/*
 * @lc app=leetcode.cn id=477 lang=golang
 *
 * [477] 汉明距离总和
 */

// @lc code=start
func totalHammingDistance(nums []int) int {
	// num < 10^9
	// 2^30 > 10^9
	num_of_1 := make([]int,31)
	n := len(nums)
	for _,num := range nums {
		for i := 0 ; num >> i != 0 ; i++ {
			if (num >> i) & 1 == 1 {
				num_of_1[i]++
			}
		}
	}
	result := 0
	for i := 0 ; i < 31 ; i++ {
		if num_of_1[i] > 0 && num_of_1[i] < n {
			result += num_of_1[i] * (n-num_of_1[i])
		}
	}
	return result
}
// @lc code=end

