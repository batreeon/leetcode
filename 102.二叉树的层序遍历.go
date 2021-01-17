/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		lastfloor := len(queue)
		for i := 0; i < lastfloor; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		eachfloor := make([]int, 0)
		for i := 0; i < lastfloor; i++ {
			eachfloor = append(eachfloor, queue[0].Val)
			queue = queue[1:]
		}
		ans = append(ans, eachfloor)
	}
	return ans
}

// @lc code=end

