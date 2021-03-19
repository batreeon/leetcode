/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	/*
	if left == right {
		return head
	}
	reverse := func(head *ListNode,left,right int) (*ListNode,*ListNode) {
		hhead := new(ListNode)
		last := head
		for left <= right {
			next := head.Next
			head.Next = hhead
			hhead = head
			head = next 
			left++
		}
		return hhead,last
	}
	pLeft := head
	for i := 1 ; i < left-1 ; i++ {
		pLeft = pLeft.Next
	}
	pRight := pLeft
	for i := left-1 ; i <= right ; i++ {
		pRight = pRight.Next
	} 
	hhead,last := reverse(pLeft.Next,left,right)
	pLeft.Next = hhead
	last.Next = pRight
	return head
	*/
	reverse := func(head *ListNode,left,right int) *ListNode {
		hhead := new(ListNode)
		next := new(ListNode)
		last := head
		for left <= right {
			next = head.Next
			head.Next = hhead
			hhead = head
			head = next 
			left++
		}
		last.Next = next	//哈哈哈，这一步省了不少事，嚯嚯嚯
		return hhead
	}

	if left == 1 {
		head = reverse(head,left,right)
	}else{
		pre := head
		for i := 1 ; i < left-1 ; i++ {
			pre = pre.Next
		}
		pre.Next = reverse(pre.Next,left,right)
	}

	return head
}
// @lc code=end

