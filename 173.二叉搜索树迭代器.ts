/*
 * @lc app=leetcode.cn id=173 lang=typescript
 *
 * [173] 二叉搜索树迭代器
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
// 二叉树的中序遍历
class BSTIterator {
    nodes: TreeNode[] = new Array<TreeNode>();

    constructor(root: TreeNode | null) {
        let cur:TreeNode = root;
        for (;cur !== null;) {
            this.nodes.push(cur);
            cur = cur.left;
        }
    }

    next(): number {
        let node: TreeNode = this.nodes.pop();
        let res: number = node.val;
        if (node.right !== null) {
            node = node.right;
            for (;node !== null;) {
                this.nodes.push(node);
                node = node.left;
            }
        }
        return res;
    }

    hasNext(): boolean {
        return this.nodes.length > 0;
    }
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * var obj = new BSTIterator(root)
 * var param_1 = obj.next()
 * var param_2 = obj.hasNext()
 */
// @lc code=end

