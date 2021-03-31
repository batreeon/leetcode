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
			// 也不一定是考虑过并撤销了，也可能是压根没考虑，
			// 例如一串相同的元素中的第二个为首的情况，
			// 因为第一个为首已经考虑过了，第二个为首只会导致重复，
			// 因此压根不会再将第二个为首
			// 若当前元素在一串相同元素中但不是第一个，并且其前一个元素考虑过也还没撤销，
			// 那么就说明当前元素不会做这一串相同元素的首，可以考虑
			// 子集II中没有用map来防止重复，这里可不可以也不用map呢，
			// 但是在子集中，遍历是从start开始，start处的元素还是会加进去，
			// 可以保证相同元素在不同级层的加入
			// 但是在这里遍历是从0开始的，如果不用map，那么只有相同元素序列中首位会加入最终的结果

			// 啊啊啊，另一个角度看，map的作用不是为了防止同一层级加入相同的元素，因为这个靠下标就可以做到
			// map的作用是为了能在不同层使非首位的相同元素能够加入
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

