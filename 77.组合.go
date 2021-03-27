/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	ans := [][]int{}
	var backtracking func(res []int) 
	backtracking = func(res []int) {
		if len(res) == k {
			resCopy := make([]int,k)
			copy(resCopy,res)
			ans = append(ans,resCopy)
			return
		}
		if len(res) == 0 {
			for i := 1 ; i <= n ; i++ {
				res = append(res,i)
				backtracking(res)
				res = res[:len(res)-1]
			}
		}else{
			for i := res[len(res)-1]+1 ; i <= n ; i++ {
				res = append(res,i)
				backtracking(res)
				res = res[:len(res)-1]
			}
		}
	}
	backtracking([]int{})
	return ans
}
// @lc code=end

