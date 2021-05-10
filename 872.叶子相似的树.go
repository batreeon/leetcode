/*
 * @lc app=leetcode.cn id=872 lang=golang
 *
 * [872] 叶子相似的树
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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1 := []int{}
	nodes := []*TreeNode{}

	// 得到root1的叶值序列
	for len(nodes) > 0 || root1 != nil {
		for root1 != nil {
			nodes = append(nodes,root1)
			root1 = root1.Left
		}
		node := nodes[len(nodes)-1]
		if node.Left == nil && node.Right == nil {
			leaves1 = append(leaves1,node.Val)
		}
		nodes = nodes[:len(nodes)-1]
		root1 = node.Right
	}

	// 比对root2的叶值序列
	i := 0
	for len(nodes) > 0 || root2 != nil {
		for root2 != nil {
			nodes = append(nodes,root2)
			root2 = root2.Left
		}
		node := nodes[len(nodes)-1]
		if node.Left == nil && node.Right == nil {
			if i >= len(leaves1) || node.Val != leaves1[i] {
				return false
			}
			i++
		}
		nodes = nodes[:len(nodes)-1]
		root2 = node.Right
	}

	return i == len(leaves1)
}
// @lc code=end

