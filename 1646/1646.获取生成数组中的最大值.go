/*
 * @lc app=leetcode.cn id=1646 lang=golang
 *
 * [1646] 获取生成数组中的最大值
 */

// @lc code=start
package main
var nums = []int{0,1,1,2,1,3,2,3}
var maxNums = []int{0,1,1,2,2,3,3,3}
func getMaximumGenerated(n int) int {
	// 当使用全局的数组存结果时，注意这第一句判断，如果不够再添
	if n < len(maxNums) {return maxNums[n]}
	l := len(maxNums)
	for i := l ; i <= n ; i++ {
		if i&1 == 0 {
			numi := nums[i/2]
			nums = append(nums,numi)
			maxNums = append(maxNums,max(maxNums[i-1],numi))
		}else{
			numi := nums[i/2]+nums[i/2+1]
			nums = append(nums,numi)
			maxNums = append(maxNums,max(maxNums[i-1],numi))
		}
	}
	return maxNums[n]
}
func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

