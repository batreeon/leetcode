/*
 * @lc app=leetcode.cn id=1669 lang=golang
 *
 * [1669] 合并两个链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	p := list1
	var prea *ListNode
	for i := 0; i <= b; i++ {
		if a-1 == i {
			prea = p
		}
		p = p.Next
	}
	prea.Next = list2
	
	p2 := list2
	for p2.Next != nil {
		p2 = p2.Next
	}

	p2.Next = p
	return list1
}
// @lc code=end

