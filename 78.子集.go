/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	// 回溯
	// 最终结果
	result := [][]int{}
	// 中间结果
	res := []int{}
	var backTrack  func(nums []int , pos int , res []int , result *[][]int)
	// pos是下一个要添加到集合中的元素下标
	// res参数是将要加入到result
	backTrack = func(nums []int , pos int , res []int , result *[][]int) {
		ans := make([]int,len(res))
		copy(ans,res)
		*result = append(*result,ans)
		for i := pos ; i < len(nums) ; i++ {
			res = append(res,nums[i])
			backTrack(nums,i+1,res,result)
			res = res[:len(res)-1]
		}
	}
	backTrack(nums,0,res,&result)
	return result
}
// @lc code=end

