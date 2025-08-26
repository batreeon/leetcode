#
# @lc app=leetcode.cn id=82 lang=python3
#
# [82] 删除排序链表中的重复元素 II
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(next=head)
        pre = dummy
        cur = head
        while cur:
            if cur and cur.next:
                if cur.val == cur.next.val:
                    while cur.val == cur.next.val:
                        cur = cur.next
                        if not cur.next:
                            break
                    pre.next = cur.next
                else:
                    pre = cur
            cur = cur.next

        return dummy.next
# @lc code=end

