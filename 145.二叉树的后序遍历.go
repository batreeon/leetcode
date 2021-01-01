/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {
	ans := make([]int,0)
	stack := make([]*TreeNode,0)
	var lastNode *TreeNode
	//记录上一个出栈的节点
	for ;; {
		for ;root != nil; {
			stack = append(stack,root)
			root = root.Left
		}
		if len(stack) == 0 {
			break
		}
		node := stack[len(stack)-1]
		//拿出来但不一定出栈，必须等他的右孩子出栈后才能出栈
		if node.Right == nil || node.Right == lastNode {
			ans = append(ans,node.Val)
			stack = stack[:len(stack)-1]
			lastNode = node
		}else {
			root = node.Right
		}
	}
	return ans
}
// @lc code=end

