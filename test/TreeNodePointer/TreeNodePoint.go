package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// pNode可以用来修改root的属性，重定义一个树节点指针是为了保留root
func main() {
	root := &TreeNode{Val: 1}
	pNode := root
	pNode.Val = 1 + pNode.Val
	fmt.Println(root.Val)
}
