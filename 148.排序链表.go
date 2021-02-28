/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var mergeSortList func(*ListNode , *ListNode) *ListNode
	var merge func(*ListNode , *ListNode) *ListNode


	mergeSortList = func(head , tail *ListNode) *ListNode {//记一记
		//1->2
		//sL(1,nil) -> sl(1,2) = 1,sl(2,nil)=2
		//1->2->3
		//sL(1,nil) -> sl(1,3) = sl(1,2) = 1,sl(2,3)=2  sl(3,nil) = 3
		if head.Next == nil {
			return head
		}
		if head.Next == tail {
			head.Next = nil
			return head
		}
		//用快慢指针法找到中点
		slow,fast := head,head
		for fast != tail {
			slow = slow.Next
			fast = fast.Next
			if fast != tail {
				fast = fast.Next
			}
		}
		left := mergeSortList(head,slow)
		right := mergeSortList(slow,tail)
		sorted := merge(left,right)
		return sorted
	}
	merge = func(l,r *ListNode) *ListNode {
		hHead := &ListNode{}
		pHead := hHead
		for l != nil && r != nil {
			if l.Val < r.Val {
				pHead.Next = l 
				l = l.Next
			}else{
				pHead.Next = r
				r = r.Next
			}
			pHead = pHead.Next
		}
		if l != nil {
			pHead.Next = l
		}else{
			pHead.Next = r
		}
		return hHead.Next
	}

	return mergeSortList(head,nil)
}
// @lc code=end

