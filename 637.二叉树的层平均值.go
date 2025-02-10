/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
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
func averageOfLevels(root *TreeNode) []float64 {
    level := []*TreeNode{root}
	var res []float64
	for len(level) > 0 {
		l := len(level)
		lsum := float64(0)
		for i := 0; i < l; i++ {
			if level[i].Left != nil {
				level = append(level, level[i].Left)
			}
			if level[i].Right != nil {
				level = append(level, level[i].Right)
			}
			lsum += float64(level[i].Val)
		}
		res = append(res, lsum/float64(l))
		level = level[l:]
	}
	return res
}
// @lc code=end

