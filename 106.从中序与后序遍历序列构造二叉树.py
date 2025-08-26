#
# @lc app=leetcode.cn id=106 lang=python3
#
# [106] 从中序与后序遍历序列构造二叉树
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def buildTree(self, inorder: List[int], postorder: List[int]) -> Optional[TreeNode]:
        if not inorder:
            return None
        root_val = postorder[-1]
        i = inorder.index(root_val)
        li, ri = inorder[:i], inorder[i+1:]
        lp, rp = postorder[:i], postorder[i:-1]
        root = TreeNode(root_val)
        root.left = self.buildTree(li, lp)
        root.right = self.buildTree(ri, rp)
        return root
# @lc code=end

