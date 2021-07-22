/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	for pNode := head ; pNode != nil ; {
		newNode := &Node{Val:pNode.Val,Next:pNode.Next}
		pNode.Next = newNode
		pNode = newNode.Next
	}
	for pNode := head ; pNode != nil ; {
		if pNode.Random != nil {
			pNode.Next.Random = pNode.Random.Next
		}
		// }else{
		// 	pNode.Next.Random = pNode.Random
		// }
		pNode = pNode.Next.Next
	}
	hhead := head.Next
	for pNode := head ; pNode != nil ; {
		next := pNode.Next.Next
		if next == nil {
			pNode.Next = nil
			break
		}
		pNode.Next.Next = next.Next
		pNode.Next = next
		pNode = next
	}
	return hhead
}
// @lc code=end

