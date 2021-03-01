/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	var ans int
	for root != nil || len(stack) != 0 {//借用了树的遍历，这里是前序
		for root != nil {
			stack = append(stack,root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		k--
		if k == 0 {
			ans = node.Val
		}
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return ans
}
// @lc code=end

