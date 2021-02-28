/*
 * @lc app=leetcode.cn id=147 lang=golang
 *
 * [147] 对链表进行插入排序
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	/*
	if head == nil {
		return head
	}
	curNode := head.Next
	lastNode := head
	for curNode != nil {
		if curNode.Val <= head.Val {
			lastNode.Next = curNode.Next
			curNode.Next = head
			head = curNode
			curNode = lastNode.Next
		}else{
			pNode := head
			for curNode.Val > pNode.Next.Val {
				pNode = pNode.Next
			}
			//替代pNode
			lastNode.Next = curNode.Next
			//这里不好解决啊，如果curNode移到前面去了，lastNode只改变Next
			//如果curNode还在原位，那么下一步lastNode就等于curNode
			//这样的话用我这个判断结构就有点麻烦，只能像下面官解那样写了
			curNode.Next = pNode.Next
			pNode.Next = curNode
			curNode = lastNode.Next
		}
	}
	return head
	*/

	
	if head == nil {
		return head
	}
	//用来作为head的指引
	hHead := &ListNode{Next:head}
	//维护已排序链
	sortedTail := head
	curNode := head.Next
	for curNode != nil {
		if curNode.Val > sortedTail.Val {
			sortedTail = curNode
		}else{
			preNode := hHead
			for curNode.Val > preNode.Next.Val {
				preNode = preNode.Next
			}
			sortedTail.Next = curNode.Next
			curNode.Next = preNode.Next
			preNode.Next = curNode
		}
		curNode = sortedTail.Next
	}
	return hHead.Next
	
}
// @lc code=end

