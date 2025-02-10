/*
 * @lc app=leetcode.cn id=98 lang=typescript
 *
 * [98] 验证二叉搜索树
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

// 递归，检查二叉搜索树
function isValidBST(root: TreeNode | null): boolean {
    return dfs(root).isValid;
};

function dfs(node: TreeNode): {isValid: boolean, min: number, max: number} {
    if (node === null) {
        return {
            isValid: true,
            min: 2147483648, // 这俩极值选取得注意啊
            max: -2147483649
        }
    }
    let left = dfs(node.left);
    let right = dfs(node.right);
    if (!left.isValid || !right.isValid || left.max >= node.val || node.val >= right.min) {
        return {
            isValid: false,
            min: 0,
            max: 0
        };
    }
    return {
        isValid: true,
        min: Math.min(node.val, left.min),
        max: Math.max(node.val, right.max)
    };
}
// @lc code=end

