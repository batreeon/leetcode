#
# @lc app=leetcode.cn id=117 lang=python3
#
# [117] 填充每个节点的下一个右侧节点指针 II
#

# @lc code=start
"""
# Definition for a Node.
class Node:
    def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next
"""

class Solution:
    def connect(self, root: 'Node') -> 'Node':
        if not root:
            return None
        if root.left and root.right:
            root.left.next = root.right
        elif root.left:
            root.left.next = self.generate_next(root.next)
        if root.right:
            root.right.next = self.generate_next(root.next)

        self.connect(root.right)
        self.connect(root.left)

        return root
    
    def generate_next(self, root: 'Node') -> 'None':
        if not root:
            return None
        if root.left:
            return root.left
        if root.right:
            return root.right
        return self.generate_next(root.next)
# @lc code=end

