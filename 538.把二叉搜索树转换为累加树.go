/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
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
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	nodes := []*TreeNode{}
	stack := []*TreeNode{}
	pNode := root
	// 中序遍历
	// 在这里pNode != nil是一定要的，当将根节点退出后，satck长度为0，但其Right不为空
	for pNode != nil || len(stack) > 0 {
		for pNode != nil {
			stack = append(stack, pNode)
			pNode = pNode.Left
		}
		pNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nodes = append(nodes, pNode)
		pNode = pNode.Right
	}
	sum := 0
	for i := len(nodes) - 1; i >= 0; i-- {
		nodes[i].Val = nodes[i].Val + sum
		sum = nodes[i].Val
	}
	return root
}

// @lc code=end

