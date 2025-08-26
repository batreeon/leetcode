#
# @lc app=leetcode.cn id=105 lang=python3
#
# [105] 从前序与中序遍历序列构造二叉树
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        if not preorder:
            return None
        root_val = preorder[0]
        i = inorder.index(root_val)
        left_preorder, right_preorder = preorder[1:1+i], preorder[1+i:]
        left_inorder, right_inorder = inorder[:i], inorder[i+1:]

        root = TreeNode(root_val)
        root.left = self.buildTree(left_preorder, left_inorder)
        root.right = self.buildTree(right_preorder, right_inorder)
        return root
# @lc code=end

