/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    m := make(map[*ListNode]bool)
	// 不初始化，这样他初始状态就为nil
	var result *ListNode
	headToTail := func(root *ListNode) {
		for root != nil {
			if m[root] {
				result = root
				return
			}else{
				m[root] = true
			}
			root = root.Next
		}
		return
	}
	headToTail(headA)
	headToTail(headB)
	return result
}
// @lc code=end

