#
# @lc app=leetcode.cn id=230 lang=python3
#
# [230] 二叉搜索树中第 K 小的元素
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        nodes = []
        cur = root
        while cur or nodes:
            while cur:
                nodes.append(cur)
                cur = cur.left
            
            node = nodes.pop()
            k -= 1
            if k == 0:
                return node.val
            
            cur = node.right
# @lc code=end

