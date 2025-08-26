#
# @lc app=leetcode.cn id=124 lang=python3
#
# [124] 二叉树中的最大路径和
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        self.result = -inf
        self.max_path_sum(root)
        return self.result
    
    def max_path_sum(self, root: Optional[TreeNode]) -> int:
        if root == None:
            return 0
        
        l, r = max(0, self.max_path_sum(root.left)), max(0, self.max_path_sum(root.right))
        self.result = max(self.result , root.val + l + r)
        return root.val + max(l,r)
        
# @lc code=end

