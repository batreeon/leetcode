/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	//如果有重复元素，这个就不适用了
	n := len(nums)
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return append([][]int{},nums)
	}
	ans := [][]int{}
	for _,v := range permute(nums[1:]) {
		for i := 0 ; i < n ; i++ {
			res := make([]int,n)
			copy(res[:i],v[:i])
			copy(res[i+1:],v[i:])
			res[i] = nums[0]//每次取头元素，然后插入到后续元素的排列结果中
			ans = append(ans,res)
		} 
	}
	return ans
}
// @lc code=end

