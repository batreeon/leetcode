/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	var ans []int
	var node []*TreeNode//结点栈，当作栈使用
	if root == nil {
		return ans
	}
	nownode := root
	for ;; {
		for ;nownode != nil; {
			ans = append(ans,nownode.Val)
			if nownode.Right != nil {
				node = append(node,nownode.Right)//push
			}
			nownode = nownode.Left;
		}
		if len(node) == 0 {
			break;
		}
		nownode = node[len(node)-1]//取栈顶元素
		node = node[:len(node)-1]//pop
	}
	return ans
}
// @lc code=end

