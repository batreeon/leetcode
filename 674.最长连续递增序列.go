/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 */

// @lc code=start
func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := 1
	var nowlen int
	for i := 0;i < n-1;i++ {
		if i == 0 {
			nowlen = 1
			if nums[i] < nums[i+1] {
				nowlen++
				ans = max(ans,nowlen)
			}
		}else if nums[i] >= nums[i+1] {
			nowlen = 1
		}else{
			nowlen++
			ans = max(ans,nowlen)
		}
	}
	return ans
	
	//官方解答这个思想挺不错的，比较简单：
	// start := 0
	// ans := 0
	// for i,v := range nums{
	// 	if i > 0 && v <= nums[i-1] {
	// 		start = i
	// 	}
	// 	ans = max(ans,i-start+1)
	// }
	// return ans
}
func max(x,y int) int {
	if x > y {
		return x
	}else{
		return y
	}
}
// @lc code=end

