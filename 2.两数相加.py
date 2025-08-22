#
# @lc app=leetcode.cn id=2 lang=python3
#
# [2] 两数相加
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        result = ListNode()
        p = result
        carry = 0
        while l1 or l2:
            v1 = l1.val if l1 else 0
            v2 = l2.val if l2 else 0
            sum = v1 + v2 + carry
            carry = sum // 10
            p.next = ListNode(val = sum%10)
            p = p.next
            l1 = l1.next if l1 else l1
            l2 = l2.next if l2 else l2
        if carry > 0:
            p.next = ListNode(val = carry)
        return result.next

# @lc code=end

