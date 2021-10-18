/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	result := rootSum(root, targetSum)
	result += pathSum(root.Left, targetSum)
	result += pathSum(root.Right, targetSum)
	return result
}

// 以root为根节点是否存在路径
func rootSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	result := 0
	if root.Val == targetSum {
		result++
	}
	result += rootSum(root.Left, targetSum - root.Val)
	result += rootSum(root.Right, targetSum - root.Val)
	return result
}
// @lc code=end

