/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	lessQ , moreQ := []int{},[]int{}
	for node := head ; node != nil ; {
		if node.Val < x {
			lessQ = append(lessQ,node.Val)
		}else{
			moreQ = append(moreQ,node.Val)
		}
		node = node.Next
	}
	nowNode := head
	for _,lx := range lessQ {
		nowNode.Val = lx
		nowNode = nowNode.Next
	}
	for _,mx := range moreQ {
		nowNode.Val = mx
		nowNode = nowNode.Next
	}
	return head
}
// @lc code=end

