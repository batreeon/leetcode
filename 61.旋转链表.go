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
	/*
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
	*/

	// 20210327
	if head == nil || head.Next == nil {
		return head
	}
	
	// 找到链表长度和尾节点
	tail := head
	l := 1
	for tail.Next != nil {
		tail = tail.Next
		l++
	}

	// 有效旋转长度
	k %= l
	if k == 0 {
		return head
	}

	// 找到新的头和尾
	newTail,newHead := head,head.Next	//旋转l-1次的结果
	// 则要得到旋转k次的结果，应执行l-(l-k)次
	for i := 1 ; l-k != i ; i++ {
		newTail,newHead = newTail.Next,newHead.Next
	}

	tail.Next = head
	head = newHead
	newTail.Next = nil
	return head
}
// @lc code=end

