/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	 /*
	 //感觉没错啊，没找到错在哪
	treeNodeMap := map[int]*TreeNode{}	//用来记录非nil节点和其对应的满二叉树序号（根节点为1）
	pN,qN := p.Val,q.Val
	pI,qI := -1,-1

	treeNodeMap[1] = root
	keyPool := []int{1}
	loop:
		// for len(keyPool) > 0 {
		for {
			l := len(keyPool)
			for i := 0 ; i < l ; i++ {
				key := keyPool[i]
				node := treeNodeMap[key]
				v := node.Val
				if v == pN {
					pI = key
				}
				if v == qN {
					qI = key
				}

				if pI != -1 && qI != -1  {
					break loop
				}

				if node.Left != nil {
					treeNodeMap[2*key] = node.Left
					keyPool = append(keyPool,2*key)
				}
				if node.Right != nil {
					treeNodeMap[2*key+1] = node.Right
					keyPool = append(keyPool,2*key+1)
				}

			}
			keyPool = keyPool[l:len(keyPool)]
		}

	//找到祖先
	pI,qI = int(pI),int(qI)
	for pI != qI {
		if pI > qI {
			pI = pI / 2
		}else{
			qI = qI / 2
		}
	}
	return treeNodeMap[pI]
	*/
	
	//官解
	if root == nil || p == root || q == root {
		return root
	}
	left := lowestCommonAncestor(root.Left,p,q)
	right := lowestCommonAncestor(root.Right,p,q)
	if left == nil && right == nil {
		return nil
	}else if left != nil && right != nil {
		return root
	}else{
		if left == nil {
			return right
		}else{
			return left
		}
	}
}
// @lc code=end

