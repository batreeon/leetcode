/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根节点到叶节点数字之和
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
func sumNumbers(root *TreeNode) int {
	result := 0
	var f func(preSum int, root *TreeNode)
	f = func(preSum int, root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			result += preSum + root.Val
		} else {
			f(10*(preSum+root.Val), root.Left)
			f(10*(preSum+root.Val), root.Right)
		}
	}
	f(0, root)
	return result
}

// @lc code=end

