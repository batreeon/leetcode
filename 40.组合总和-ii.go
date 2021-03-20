/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
package main
import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	l := len(candidates)
	sort.Ints(candidates)

	ans := [][]int{}
	visited := map[int]bool{}
	// 用指针做参数，减小拷贝消耗
	// 取得值必须是递增的，所以只有满足了target - candidates[i] >= candidates[minI]才继续进行
	var cs func(res *[]int, target int, minI int)
	cs = func(res *[]int, target int, minI int) {
		for i := minI + 1; i < l; i++ {
			if i != 0 && candidates[i-1] == candidates[i] && !visited[i-1] {
				continue
			}
			if target-candidates[i] == 0 {
				*res = append(*res, candidates[i])

				// 这一步很重要，不能直接传*res
				ans1 := make([]int, len(*res))
				copy(ans1, *res)
				ans = append(ans, ans1)

				*res = (*res)[:len(*res)-1]
			//下面这个判断条件是可以让程序提早退出的，减少了时间的浪费 
			} else if minI == -1 || (minI >= 0 && target-candidates[i] >= candidates[minI]) {
			// 若直接使用else,也可以得到答案，但一些用例会超时
			// }else{
				*res = append(*res, candidates[i])
				visited[i] = true
				cs(res, target-candidates[i], i)
				*res = (*res)[:len(*res)-1]
				delete(visited,i)
			}
		}
	}
	cs(&[]int{}, target, -1)
	return ans
}
// @lc code=end

