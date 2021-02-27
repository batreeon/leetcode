/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	// 有点慢
    // nodeMap := map[*ListNode]bool{}
	// pNode := head 
	// for pNode != nil {
	// 	if nodeMap[pNode] == true {
	// 		return pNode
	// 	}
	// 	nodeMap[pNode] = true
	// 	pNode = pNode.Next
	// }
	// return nil

	// 快慢指针法
	fast,slow := head,head
	for fast != nil {
		slow = slow.Next 
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next 
		if fast == slow {//相遇了
			//起始点到入环点的距离 = f,s相遇点到入环点距离+n*环长
			//因此最终p和slow一定能相遇
			p := head
			for p != slow {
				p = p.Next 
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
// @lc code=end

