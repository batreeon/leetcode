/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	var backTrack func(nums []int , pos int , res []int , result *[][]int)
	backTrack = func(nums []int , pos int , res []int , result *[][]int) {
		ans := make([]int,len(res))
		copy(ans,res)
		*result = append(*result,ans)
		for i := pos ; i < len(nums) ; i++ {
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
}
// @lc code=end

