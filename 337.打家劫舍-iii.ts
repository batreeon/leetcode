/*
 * @lc app=leetcode.cn id=337 lang=typescript
 *
 * [337] 打家劫舍 III
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

function rob(root: TreeNode | null): number {
    return Math.max(...dfs(root));
};

function dfs(root: TreeNode | null): [number, number] {
    if (root === null) {
        return [0,0];
    }
    let [lroot, lnoroot] = dfs(root.left);
    let [rroot, rnoroot] = dfs(root.right);
    return [root.val + lnoroot + rnoroot, Math.max(lroot, lnoroot) + Math.max(rroot, rnoroot)]; 
}
// @lc code=end

