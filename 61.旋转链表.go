/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	pNode := head
	var tail *ListNode
	l := 0
	for pNode != nil {
		l++
		tail = pNode
		pNode = pNode.Next
	}//得到尾节点和链表长度

	k = k%l //得到有效旋转次数
	if l == 1 || k == 0 {
		return head
	}//不用旋转，直接返回

	newTail := head
	pNode = head.Next
	for i := 1 ; pNode != nil ; i++{
		if l - k == i {
			break
		}
		newTail = newTail.Next
		pNode = pNode.Next
	}//找到新的头和尾

	tail.Next = head
	head = pNode
	newTail.Next = nil//修改相应的量
	return head
}
// @lc code=end

