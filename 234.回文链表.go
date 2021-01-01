/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	//len := 0
	var s []int
	for p := head; p != nil;p=p.Next {
		s = append(s,p.Val)
	}
	len := len(s)
	for begin := 0;begin < len/2;begin++ {
		if s[begin] != s[len-begin-1] {
			return false
		}
	}
	return true
}
// @lc code=end

