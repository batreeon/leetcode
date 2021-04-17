/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
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
func rob(root *TreeNode) int {
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	
	var robroot func(root *TreeNode) (int,int) 
	robroot = func(root *TreeNode) (int,int) {
		// 返回值分别是包含root自身和不包含root自身得到的最大偷盗值
		if root.Left == nil && root.Right == nil {
			return root.Val,0
		}
		leftwithroot,leftwithoutroot := 0,0
		rightwithroot,rightwithoutroot := 0,0
		if root.Left != nil {
			leftwithroot,leftwithoutroot = robroot(root.Left)
		}
		if root.Right != nil {
			rightwithroot,rightwithoutroot = robroot(root.Right)
		}
		return root.Val+leftwithoutroot+rightwithoutroot,max(leftwithroot,leftwithoutroot)+max(rightwithroot,rightwithoutroot)
	}

	return max(robroot(root))
}
// @lc code=end
