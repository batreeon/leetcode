/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
package main
import "sort"


func permuteUnique(nums []int) [][]int {
	// 有重复元素，相同值同一级只能选一次
	// 如何判断是否同一级是否选过了呢
	// 如果当前数不是当前相同数的第一个，那么如果前一个数visited了，那么当前元素就可以选，说明是下一级了
	// 若前一个相同元素，非visited，那么说明前一个已经选过并且撤销了，当前元素就不需要再考虑了
	result := [][]int{}
	res := []int{}
	sortedNums := make([]int,len(nums))
	copy(sortedNums,nums)
	sort.Ints(sortedNums)
	visited := map[int]bool{}
	n := len(nums)

	var backTrack func(nums []int,res []int,result *[][]int)
	backTrack = func(nums []int,res []int,result *[][]int) {
		if len(res) == n {
			resCopy := make([]int,len(res))
			copy(resCopy,res)
			*result = append(*result,resCopy)
			return
		}
		for i := 0 ; i < len(nums) ; i++ {
			//避免重复添加
			if visited[i] {
				continue
			}

			// 前两个条件，确保不是相同元素中的第一个
			// 第三个条件，说明相同元素已经考虑过并且撤销了，并且当前元素不是前一个元素的下一级
			if i != 0 && nums[i] == nums[i-1] && !visited[i-1] {
				continue
			}
			res = append(res,nums[i])
			visited[i] = true
			backTrack(nums,res,result)
			res = res[:len(res)-1]
			delete(visited,i)
		}
	}
	backTrack(sortedNums,res,&result)
	return result
}
// @lc code=end

