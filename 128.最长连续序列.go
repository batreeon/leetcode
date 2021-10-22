/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for _, num := range nums {
		m[num] = true
	}
	result := 0
	for k := range m {
		// 如果k-1存在，那么以k开头的序列一定不是最终答案，就不需要浪费时间了
		if !m[k-1] {
			curNum := k
			curLen := 1
			for m[curNum+1] {
				curNum++
				curLen++
			}
			if curLen > result {
				result = curLen
			}
		}
	}
	return result
}
// @lc code=end

