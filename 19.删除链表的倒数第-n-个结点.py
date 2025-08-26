#
# @lc app=leetcode.cn id=19 lang=python3
#
# [19] 删除链表的倒数第 N 个结点
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy = ListNode(next=head)
        pre = dummy
        cur = head
        for _ in range(n):
            cur = cur.next
        
        while cur:
            pre,cur = pre.next, cur.next

        pre.next = pre.next.next

        return dummy.next
# @lc code=end

