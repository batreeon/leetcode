/*
 * @lc app=leetcode.cn id=229 lang=golang
 *
 * [229] 求众数 II
 */

// @lc code=start
func majorityElement(nums []int) []int {
	ans := []int{}
	n := len(nums)
	// n<3,则n/3=0，则在下方第二个循环中，不会有结果进入最终的ans
	// 因此单独再走一个分支
	if n < 3 {
		m := map[int]bool{}
		for _,v := range nums {
			if m[v] == false {
				ans = append(ans,v)
				m[v] = true
			}
		}
		return ans
	}
	m := map[int]int{}
	n /= 3
	for _,v := range nums {
		if times,ok := m[v] ; !ok {
			m[v]++
		}else if times <= n {
			m[v]++
			if m[v] > n {
				ans = append(ans,v)
			}
		}
	}
	return ans
}
// @lc code=end

