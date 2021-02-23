/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for _,v := range nums {
		x := (v-1)%n 
		//下标是0~n-1,数的范围1~n，看似不用取模，但是注意nums[x] += n,经过这个就越界了，需要取模
		nums[x] += n
	}
	ans := []int{}
	for i,v := range nums {
		if v < n+1 {
			ans = append(ans,i+1)
		}
	}
	return ans
}
// @lc code=end

