/*
 * @lc app=leetcode.cn id=236 lang=typescript
 *
 * [236] 二叉树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

function lowestCommonAncestor(root: TreeNode | null, p: TreeNode | null, q: TreeNode | null): TreeNode | null {
	if (root === null || root === p || root === q) {
        return root;
    }
    let l: TreeNode = lowestCommonAncestor(root.left, p, q);
    let r: TreeNode = lowestCommonAncestor(root.right, p, q);
    if (!l && !r) {
        return null;
    }
    if (l && r) {
        return root;
    }
    return l ? l : r;
};
// @lc code=end

