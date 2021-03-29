/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
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

func getNext(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		return root.Left
	}else if root.Right != nil {
		return root.Right
	}else{
		return getNext(root.Next)
	}
} 
func connect(root *Node) *Node {

	/*
	// 层次遍历，空间复杂度不符合
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
	*/

	// 要处理完一层再处理下一层
	if root == nil || (root.Left == nil && root.Right == nil){
		return root
	}
	if root.Left != nil && root.Right != nil {
		root.Left.Next = root.Right
	}else if root.Right == nil {
		root.Left.Next = getNext(root.Next)
	}
	if root.Right != nil {
		root.Right.Next = getNext(root.Next)
	}
	// 下面这两句不能颠倒顺序，因为connect(root.Left)时，可能右侧的还没有进行处理，而connect(root.Right)时不会用到左侧的
	// 比如用力[2,1,3,0,7,9,1,2,null,1,0,null,null,8,8,null,null,null,null,7]，answer:[2,#,1,3,#,0,7,9,1,#,2,1,0,8,8,#,7,#]
	// 若先处理left，则会丢失8,8
	connect(root.Right)
	connect(root.Left)
	return root
}
// @lc code=end

