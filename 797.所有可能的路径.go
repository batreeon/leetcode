/*
 * @lc app=leetcode.cn id=797 lang=golang
 *
 * [797] 所有可能的路径
 */

// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {

	/*
	n := len(graph)
	// 记录经该点是否能到终点,不能到就记录下来,以后遇到就直接不走了
	// 但加了这个成绩没大差别
	unavailable := map[int]bool{}
	// 保存可行路径
	result := [][]int{}

	// 其实就是dfs吧
	var backtrack func(res []int) bool
	backtrack = func(res []int) bool {
		now := res[len(res)-1]
		if now == n-1 {
			resCopy := make([]int,len(res))
			copy(resCopy,res)
			result = append(result,resCopy)
			// 只有最后到达n-1了才会返回true
			return true
		}
		flag := false
		for _,next := range graph[now] {
			if unavailable[next] {
				continue
			}
			res = append(res,next)
			if backtrack(res) {
				flag = true
			}else{
				// 从next无法到n-1
				unavailable[next] = true
			}
			res = res[:len(res)-1]
		}
		return flag
	}

	backtrack([]int{0})
	return result
	*/
	
	
	// 题目是有向无环图,则可以用递归
	n := len(graph)
	// path[i]用来记录点i到n-1的所有路径，避免重复计算
	paths := make([][][]int,n)
	paths[n-1] = [][]int{[]int{n-1}}

	// recur(v)返回所有从v到n-1的路径
	var recur func(v int) [][]int 
	recur = func(v int) [][]int {
		res := [][]int{}
		for _,next := range graph[v] {
			if len(paths[next]) != 0 {
				for _,path := range paths[next] {
					res = append(res,append([]int{v},path...))
				}
			}else{
				for _,path := range recur(next) {
					res = append(res,append([]int{v},path...))
				}
			}
		}
		paths[v] = res
		return res
	}

	return recur(0)
}
// @lc code=end

