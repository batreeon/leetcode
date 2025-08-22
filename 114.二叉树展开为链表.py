#
# @lc app=leetcode.cn id=114 lang=python3
#
# [114] 二叉树展开为链表
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def flatten(self, root: Optional[TreeNode]) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        if not root:
            return None
        right = root.right
        root.right = self.flatten(root.left)
        root.left = None
        p_root = root
        while p_root.right:
            p_root = p_root.right
        p_root.right = self.flatten(right)
        return root
# @lc code=end

