#
# @lc app=leetcode.cn id=98 lang=python3
#
# [98] 验证二叉搜索树
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        r, _, _ = self.is_valid_bst(root)
        return r
    
    def is_valid_bst(self, root: Optional[TreeNode]) -> (bool, int, int):
        if root is None:
            return True, inf, -inf
        v = root.val
        l, lmin, lmax = self.is_valid_bst(root.left)
        r, rmin, rmax = self.is_valid_bst(root.right)
        valid = l and r and v > lmax and v < rmin

        return valid, min(root.val, lmin, rmin), max(root.val, lmax, rmax)
# @lc code=end

