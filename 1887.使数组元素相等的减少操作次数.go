/*
 * @lc app=leetcode.cn id=1887 lang=golang
 *
 * [1887] 使数组元素相等的减少操作次数
 */

// @lc code=start
func reductionOperations(nums []int) int {
	sort.Ints(nums)
	numss := []int{}
	for i := 0 ; i < len(nums) ; {
		begin := i
		// 初始条件i=i+1,如果有相同序列，begin记录的是起始点，i从第二个开始考察
		for i = i+1 ; i < len(nums) && nums[i] == nums[i-1] ; i++ {}
		numss = append(numss,i-begin)
	}
	result := 0
	for i := 1 ; i < len(numss) ; i++ {
		result += i*numss[i]
	}
	return result
}
// @lc code=end

