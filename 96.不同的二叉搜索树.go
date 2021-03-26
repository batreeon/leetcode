/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
func numTrees(n int) int {
	nums := make([]int,n+1)
	nums[0] = 1
	nums[1] = 1
	var numtrees func(n int) int
	numtrees = func(n int) int {
		if nums[n] != 0 {
			return nums[n]
		}
		for i := 0 ; i < n ; i++ {
			// 除去根节点，左右分别有i,n-i-1个节点
			nums[n] += numtrees(i)*numtrees(n-i-1)
		}
		return nums[n]
	}
	return numtrees(n)
}
// @lc code=end

