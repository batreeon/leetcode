/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	presum := make([]int, len(nums)+1)
	for i, num := range nums {
		presum[i+1] += presum[i] + num
	}

	// presum[i] 是[0...i-1]数的和
	// 区间[i,j-1]的和 = presum[j] - presum[i]
	//  presum[j] - presum[i] = k
	// 那么 presum[i] = presum[j] - k
	// 那么当j固定时，就要计算j前面有多少个前缀和的值=presum[j] - k
	res := 0
	m := map[int]int{}
	for _, presumj := range presum {
		res += m[presumj-k]
		m[presumj]++
	}

	return res
}
// @lc code=end

