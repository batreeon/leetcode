/*
 * @lc app=leetcode.cn id=783 lang=golang
 *
 * [783] 二叉搜索树节点最小距离
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
 package main
 import "math"
func minDiffInBST(root *TreeNode) int {
	result := math.MaxInt64
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	var mindiffinbst func(root *TreeNode) (int,int)
	mindiffinbst = func(root *TreeNode) (int,int) {
		if root.Left == nil && root.Right == nil {
			return root.Val,root.Val
		}
		var leftmin,leftmax,rightmin,rightmax int
		if root.Left != nil {
			leftmin,leftmax = mindiffinbst(root.Left)
			result = min(result,root.Val-leftmax)
		}
		if root.Right != nil {
			rightmin,rightmax = mindiffinbst(root.Right)
			result = min(result,rightmin-root.Val)
		}
		if root.Left == nil {
			leftmin = rightmin
		}
		if root.Right == nil {
			rightmax = leftmax
		}
		return min(root.Val,leftmin),max(rightmax,root.Val)
	}

	mindiffinbst(root)
	return result
}
// @lc code=end

