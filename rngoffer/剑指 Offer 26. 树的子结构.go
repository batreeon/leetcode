/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 package main
 func isSubStructure(A *TreeNode, B *TreeNode) bool {
    var compare func(a,b *TreeNode) bool
    compare = func(a,b *TreeNode) bool {
        if b == nil {
            return true
        }else if a == nil {
            return false
        }
        if a.Val != b.Val {
            return false
        }
        return compare(a.Left,b.Left) && compare(a.Right,b.Right)
    }
    
	if A != nil && B != nil {
		return compare(A,B) || isSubStructure(A.Left,B) || isSubStructure(A.Right,B)
	}
	return false
}