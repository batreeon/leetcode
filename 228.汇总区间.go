/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] 汇总区间
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	ans := []string{}
	for i := 0 ; i < len(nums) ; {
		start := i
		i++
		for i < len(nums) && nums[i]-nums[i-1]==1 {
			i++
		}

		if nums[i-1] == nums[start] {
			ans = append(ans,strconv.Itoa(nums[start]))
		}else{
			startN := strconv.Itoa(nums[start])
			endN := strconv.Itoa(nums[i-1])
			ans = append(ans,startN+string("->")+endN)
		}
	}
	return ans
}
// @lc code=end

