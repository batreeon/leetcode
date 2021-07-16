/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    nodes := []*TreeNode{root}
    l2r := true
    result := [][]int{}
    for len(nodes) > 0 {
        l := len(nodes)
        temp := make([]int,l)
        for i,node := range nodes {
            if l2r {
                temp[i] = node.Val
            }else{
                temp[l-i-1] = node.Val
            }
            if node.Left != nil {
                nodes = append(nodes,node.Left)
            }
            if node.Right != nil {
                nodes = append(nodes,node.Right)
            }
        }
        result = append(result,temp)
        nodes = nodes[l:]
        l2r = !l2r
    }
    return result
}