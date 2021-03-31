/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
package main
import "sort"

func subsetsWithDup(nums []int) [][]int {
	
	/*
	// 以前通过的版本
	var backTrack func(nums []int , pos int , res []int , result *[][]int)
	backTrack = func(nums []int , pos int , res []int , result *[][]int) {
		ans := make([]int,len(res))
		copy(ans,res)
		*result = append(*result,ans)
		for i := pos ; i < len(nums) ; i++ {
			// 之前自己写的现在却看不懂了，为什么判断重复不需用map呢
			// 哦哦哦，我懂了，这个题是求组合，因此每次遍历的下标都是从上一个字符的后面开始的
			// 那么如果遍历到的某一个元素在一串相同的数字中但不是第一个，那么在这一级中，这一相同的数字就已经被访问过了
			// 况且，即使用了map,!m[i-1]是始终为true的，因为前面的元素被访问过又丢弃了
			if i != pos && nums[i] == nums[i-1] {//防止重复
				continue
			}
			res = append(res,nums[i])
			backTrack(nums,i+1,res,result)
			res = res[:len(res)-1]
		}
	}

	sortedNums := make([]int,len(nums))
	copy(sortedNums,nums)
	sort.Ints(sortedNums)
	result := [][]int{}
	res := []int{}
	backTrack(sortedNums,0,res,&result)
	return result
	*/
	

	//20210331,每日一题
	result := [][]int{}
	n := len(nums)
	var backtrack func(res []int,start int)
	backtrack = func(res []int,start int) {
		resCopy := make([]int,len(res))
		copy(resCopy,res)
		result = append(result,resCopy)
		for i := start ; i < n ; i++ {
			if i != start && nums[i] == nums[i-1] {
				continue
			}
			res = append(res,nums[i])
			backtrack(res,i+1)
			res = res[:len(res)-1]
		}
	}
	
	sort.Ints(nums)
	backtrack([]int{},0)
	return result
}
// @lc code=end

