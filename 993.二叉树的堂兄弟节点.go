/*
 * @lc app=leetcode.cn id=993 lang=golang
 *
 * [993] 二叉树的堂兄弟节点
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
func isCousins(root *TreeNode, x int, y int) bool {
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		l := len(nodes)
		xseen,yseen := false,false
		for _,node := range nodes {
			if node.Left != nil && node.Right != nil {
				if (node.Left.Val == x && node.Right.Val == y) || (node.Left.Val == y && node.Right.Val == x) {
					return false
				}
				if node.Left.Val == x || node.Right.Val == x {
					xseen = true
				}
				if node.Left.Val == y || node.Right.Val == y {
					yseen = true
				}
				nodes = append(nodes,node.Left)
				nodes = append(nodes,node.Right)
			}
			if node.Left != nil {
				if node.Left.Val == x {
					xseen = true
				}
				if node.Left.Val == y {
					yseen = true
				}
				nodes = append(nodes,node.Left)
			}
			if node.Right != nil {
				if node.Right.Val == x {
					xseen = true
				}
				if node.Right.Val == y {
					yseen = true
				}
				nodes = append(nodes,node.Right)
			}
		}
		if xseen && yseen {
			return true
		}
		if xseen {
			return false
		}
		if yseen {
			return false
		}
		nodes = nodes[l:]
	}
	return false
}
// @lc code=end

