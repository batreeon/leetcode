/*
 * @lc app=leetcode.cn id=237 lang=golang
 *
 * [237] 删除链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	curNode := node
    for curNode.Next.Next != nil {
		curNode.Val = curNode.Next.Val
		curNode = curNode.Next
	}
	curNode.Val = curNode.Next.Val
	curNode.Next = nil
	/*
	为什么这样写不对啊
    for node.Next != nil {
		node.Val = node.Next.Val
		node = node.Next
	}
	node = nil
	*/
}
// @lc code=end

