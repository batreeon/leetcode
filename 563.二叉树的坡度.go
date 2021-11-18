/*
 * @lc app=leetcode.cn id=563 lang=golang
 *
 * [563] 二叉树的坡度
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
func findTilt(root *TreeNode) int {
	result := 0
	findtilt(root, &result)
	return result
}

// 用来返回以root为根的子树和
// 并且没计算一个根节点的坡度，就把他加到result上
func (root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	l,r := findtilt(root.Left, result), findtilt(root.Right, result)
	if l > r {
		(*result) += (l - r)
	}else{
		(*result) += (r - l)
	}

	return l + r + root.Val
}
// @lc code=end

