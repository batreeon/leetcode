/*
 * @lc app=leetcode.cn id=797 lang=golang
 *
 * [797] 所有可能的路径
 */

// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	// 当前遍历路径下是否已经走过了
	seen := make([]bool,n)
	// 记录经该点是否能到终点,不能到就记录下来,以后遇到就直接不走了
	// 但加了这个成绩没差别
	available := map[int]bool{}
	// 保存可行路径
	result := [][]int{}

	var backtrack func(res []int) bool
	backtrack = func(res []int) bool {
		now := res[len(res)-1]
		if now == n-1 {
			resCopy := make([]int,len(res))
			copy(resCopy,res)
			result = append(result,resCopy)
			return true
		}
		flag := false
		for _,next := range graph[now] {
			if !seen[next] {
				if _,ok := available[next] ; ok {
					continue
				}
				seen[next] = true
				res = append(res,next)
				if backtrack(res) {
					flag = true
				}else{
					available[next] = false
				}
				res = res[:len(res)-1]
				seen[next] = false
			}
		}
		return flag
	}

	backtrack([]int{0})
	return result
}
// @lc code=end

