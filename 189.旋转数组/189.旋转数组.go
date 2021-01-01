/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int)  {
	k %= len(nums)//防止出现len(nums) < k的情形
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])	
}

func reverse(s []int) {
	for i,j := 0,len(s)-1;i < j;i,j = i+1,j-1 {
		s[i],s[j] = s[j],s[i]
	}
}
// @lc code=end

