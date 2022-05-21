/*
 * @lc app=leetcode.cn id=961 lang=golang
 *
 * [961] 在长度 2N 的数组中找出重复 N 次的元素
 */

// @lc code=start
// 共2n个数，n+1个不同的，有一个重复了n次，那么就只有一个出现了一次以上呗
func repeatedNTimes(nums []int) int {
	count := map[int]int{}
	for _, num := range nums {
		if count[num] == 1 {
			return num
		}
		count[num]++
	}
	return -1
}
// @lc code=end

