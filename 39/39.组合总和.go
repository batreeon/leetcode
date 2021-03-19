/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	l := len(candidates)
	sort.Ints(candidates)

	ans := [][]int{}

	// 用指针做参数，减小拷贝消耗
	// 取得值必须是递增的，所以只有满足了target - candidates[i] >= candidates[minI]才继续进行
	var cs func(res *[]int, target int, minI int)
	cs = func(res *[]int, target int, minI int) {
		for i := minI; i < l; i++ {
			if target-candidates[i] == 0 {
				*res = append(*res, candidates[i])

				// 这一步很重要，不能直接传*res
				ans1 := make([]int, len(*res))
				copy(ans1, *res)
				ans = append(ans, ans1)

				*res = (*res)[:len(*res)-1]
			} else if target-candidates[i] >= candidates[minI] {
				*res = append(*res, candidates[i])
				cs(res, target-candidates[i], i)
				*res = (*res)[:len(*res)-1]
			}
		}
	}
	cs(&[]int{}, target, 0)
	return ans
}

// @lc code=end
