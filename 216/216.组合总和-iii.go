/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */

// @lc code=start
package main
func combinationSum3(k int, n int) [][]int {
	result := [][]int{}

	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	var backtrack func(pos int,res []int,goal int)
	backtrack = func(pos int,res []int,goal int) {
		// 还差一个元素
		if len(res) == k-1 {
			// goal值要能取的到
			if goal >= pos && goal <= min(n,9) {
				res = append(res,goal)
				resCopy := make([]int,k)
				copy(resCopy,res)
				result = append(result,resCopy)
			}
			return
		}

		// 还需要至少两个元素
		for i := pos ; i <= min(n,9) ; i++ {
			// 后面肯定不满足了
			if i >= goal {
				return
			}
			res = append(res,i)
			backtrack(i+1,res,goal-i)
			res = res[:len(res)-1]
		}
	}

	backtrack(1,[]int{},n)

	return result
}
// @lc code=end

