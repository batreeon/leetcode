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

	/*
	// 不符合复杂度要求吗
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
	*/

	/*
	// 大佬的解答，我去
	if root == nil {
		return root
	}
	left := root.Left
	right := root.Right
	for left != nil {
		left.Next = right 
		left = left.Right 
		right = right.Left 
	}
	connect(root.Left)
	connect(root.Right)
	return root
	*/

	// 递归解，根节点的左儿子的Next指向右儿子
	// 但右儿子的Next指向根节点的Next的左儿子
	if root == nil || (root.Left == nil && root.Right == nil){
		return root
	}
	root.Left.Next = root.Right 
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}
	connect(root.Left)
	connect(root.Right)
	return root
}
// @lc code=end

