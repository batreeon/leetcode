/*
 * @lc app=leetcode.cn id=1022 lang=golang
 *
 * [1022] 从根到叶的二进制数之和
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
func sumRootToLeaf(root *TreeNode) int {
	result := 0
	sumroottoleaf(root, 0, &result)
	return result
}

func sumroottoleaf(root *TreeNode, pre int, result *int) {
	if root == nil {
		return
	}
	pre <<= 1
	pre += root.Val
	if root.Left == nil && root.Right == nil {
		// 该节点是叶子节点
		*result += pre
		return
	}
	sumroottoleaf(root.Left, pre, result)
	sumroottoleaf(root.Right, pre, result)
}
// @lc code=end

