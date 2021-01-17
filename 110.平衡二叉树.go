/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(depth(root.Left),depth(root.Right))
}
func isBalancedRoot(root *TreeNode) bool {
	return math.Abs(float64(depth(root.Left)-depth(root.Right))) < 2
}
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalancedRoot(root) && isBalanced(root.Left) && isBalanced(root.Right)
}
// @lc code=end

