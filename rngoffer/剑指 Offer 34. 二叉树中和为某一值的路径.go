/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func pathSum(root *TreeNode, target int) [][]int {
    if root == nil {
        return [][]int{}
    }
    result := [][]int{}
    var dfs func(temp []int,r *TreeNode,presum int)
    dfs = func(temp []int,r *TreeNode,presum int) {
        presum += r.Val
        temp = append(temp,r.Val)
        if r.Left == nil && r.Right == nil && presum == target {
            tempCopy := make([]int,len(temp))
            copy(tempCopy,temp)
            result = append(result,tempCopy)
            return
        }
        if r.Left != nil {
            dfs(temp,r.Left,presum)
        }
        if r.Right != nil {
            dfs(temp,r.Right,presum)
        }
    }
    dfs([]int{},root,0)
    return result
}