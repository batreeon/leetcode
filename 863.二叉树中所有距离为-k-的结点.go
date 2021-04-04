/*
 * @lc app=leetcode.cn id=863 lang=golang
 *
 * [863] 二叉树中所有距离为 K 的结点
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
func distanceK(root *TreeNode, target *TreeNode, K int) []int {

	/*
	var findparent func(root *TreeNode) *TreeNode 
	findparent = func(root *TreeNode) *TreeNode {
		if root.Left == target {
			root.Left = nil
			return root
		}
		if root.Right == target {
			root.Right = nil
			return root
		}
		if root.Left != nil {
			left := findparent(root.Left)
			if left != nil {
				return left
			}
		}
		if root.Right != nil {
			right := findparent(root.Right)
			if right != nil {
				return right
			}
		}
		return nil
	}
	result := []int{}
	var dfs func(root *TreeNode,index int) 
	dfs = func(root *TreeNode,index int) {
		if root == nil {
			return
		}
		if index == K {
			result = append(result,root.Val)
			return
		}
		dfs(root.Left,index+1)
		dfs(root.Right,index+1)
	}

	if root == target {
		dfs(root,0)
	}else{
		dfs(target,0)
		parent := findparent(root)
		dfs(parent,1)
	}
	return result
	*/


	// 先找到每个节点的父亲,然后每次向四周发散,
	// 向外生长,队列中的元素都是距离target一样远的
	// 当距离达到K时,取出当时队列中的元素即可

	parent := map[*TreeNode]*TreeNode{}
	var dfs func(root *TreeNode,par *TreeNode) 
	dfs = func(root *TreeNode,par *TreeNode) {
		if root != nil {
			parent[root] = par
			dfs(root.Left,root)
			dfs(root.Right,root)
		}
	}
	// 从根节点开始遍历,记下每个节点的根节点
	dfs(root,nil)

	seen := map[*TreeNode]bool{}
	queue := []*TreeNode{}
	// 以target为起点,其实距离标为0,一层一层,有点像广度优先
	// 先遍历完距离为1的,再遍历距离为2的,依次,知道找的距离为K的
	queue = append(queue,target)
	seen[target] = true
	dis := 0

	for len(queue) > 0 {
		// dis==K,此时queue中存的节点就是要找的节点,取出值返回即可
		if dis == K {
			result := []int{}
			for _,node := range queue {
				result = append(result,node.Val)
			}
			return result
		}

		// 只遍历该层的节点,因此记下queue长度
		l := len(queue)
		for _,node := range queue[:l] {
			if !seen[parent[node]] && parent[node] != nil {
				queue = append(queue,parent[node])
				seen[parent[node]] = true
			}
			if !seen[node.Left] && node.Left != nil {
				queue = append(queue,node.Left)
				seen[node.Left] = true
			}
			if !seen[node.Right] && node.Right != nil {
				queue = append(queue,node.Right)
				seen[node.Right] = true
			}
		}
		// 删除上一层节点
		// 上一层并不是说上一个高度,而是距离target更近1的那一层节点
		queue = queue[l:]
		// 这一层的节点距target的距离又比上一层大1
		dis++
	}
	return []int{}
}
// @lc code=end

