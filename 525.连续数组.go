/*
 * @lc app=leetcode.cn id=525 lang=golang
 *
 * [525] 连续数组
 */

// @lc code=start
func findMaxLength(nums []int) int {
	sum := make([]int,len(nums)+1)
	for i := 0 ; i < len(nums) ; i++ {
		// 将0转化为-1，那么就是求最长的和为0的子数组
		num := -1
		if nums[i] == 1 {
			num = 1
		}
		sum[i+1] = sum[i] + num
	}
	result := 0
	// 用哈希表记录前缀和出现的最小下标
	m := make(map[int]int)
	m[0] = 0
	for i := 2 ; i < len(sum) ; i++ {
		if _,ok := m[sum[i-2]] ; !ok {
			// 只记录最早出现的，也就是最小下标
			m[sum[i-2]] = i-2
		}
		if _,ok := m[sum[i]] ; ok {
			// 先确保当前右边界下，能找到和为0的子数组
			result = max(result,i-m[sum[i]])
		}
	}
	return result
}
func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

