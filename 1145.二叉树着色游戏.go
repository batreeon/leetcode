/*
 * @lc app=leetcode.cn id=1145 lang=golang
 *
 * [1145] 二叉树着色游戏
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
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	if root.Val == x {
		left, right := dfs(root.Left), dfs(root.Right)
		if left != right {
			return true
		}
		return false
	}

	nodes := []*TreeNode{root}
	subleft, subright := 0, 0
	for len(nodes) != 0 {
		l := len(nodes)
		for i := 0; i < l; i++ {
			if nodes[i].Val == x {
				subleft = dfs(nodes[i].Left)
				subright = dfs(nodes[i].Right)
				nodes = nil
				break
			}
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
		}
		if nodes != nil {
			nodes = nodes[l:]
		}
	}

	s, m, b := sortInts([]int{subleft, subright, n - 1 - subleft - subright})
	if b > m+s {
		return true
	}
	return false
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + dfs(root.Left) + dfs(root.Right)
}

func sortInts(ints []int) (int,int,int) {
	sort.Ints(ints)
	return ints[0], ints[1], ints[2]
}
// @lc code=end

