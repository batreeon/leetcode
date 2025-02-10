/*
 * @lc app=leetcode.cn id=230 lang=typescript
 *
 * [230] 二叉搜索树中第 K 小的元素
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
//  中序遍历
function kthSmallest(root: TreeNode | null, k: number): number {
    let q: TreeNode[] = new Array<TreeNode>();
    let res: number = 0;
    for (;root !== null || q.length !==0;) {
        for (;root != null;) {
            q.push(root);
            root = root.left;
        }
        let node = q.pop();
        k--;
        if (k === 0) {
            return node.val;
        }
        root = node.right;
    }
};
// @lc code=end

