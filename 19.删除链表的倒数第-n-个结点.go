/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	hhead := new(ListNode)
	hhead.Next = head
	pNode := head
	for i := 1 ; i < n ; i++ {
		pNode = pNode.Next
	}
	for pNode.Next != nil {
		pNode = pNode.Next
		hhead = hhead.Next
	}
	if hhead.Next == head {
		head = head.Next
	}else{
		hhead.Next = hhead.Next.Next
	}
	return head
}
// @lc code=end

