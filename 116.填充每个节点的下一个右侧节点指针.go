/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	q := []*Node{}
	q = append(q,root)
	for len(q) > 0 {
		l := len(q)
		for i := 0 ; i < l ; i++ {
			if q[i].Right != nil {
				q= append(q,q[i].Right)
			}
			if q[i].Left != nil {
				q = append(q,q[i].Left)
			}
		}
		for i := 1 ; i < l ; i++ {
			q[i].Next = q[i-1]
		}
		q = q[l:]
	}
	return root

	// 当然也还可以用递归
}
// @lc code=end

