/*
 * @lc app=leetcode.cn id=671 lang=golang
 *
 * [671] 二叉树中第二小的节点
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
func findSecondMinimumValue(root *TreeNode) int {
	result := -1
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	var dfs func(r *TreeNode,rootV int)
	dfs = func(r *TreeNode,rootV int) {
		if r == nil {
			retrun
		}
		if r.Val > rootV {
			if result == -1 {
				result = r.Val
			}else{
				result = min(result,r.Val)
			}
			return
		}
		dfs(r.Left,rootV)
		dfs(r.Right,rootV)
	}
	dfs(root.Left,root.Val)
	dfs(root.Right,root.Val)
	return result
}
// @lc code=end

