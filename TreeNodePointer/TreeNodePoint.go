package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}
	pNode := root
	pNode.Val = 1 + pNode.Val
	fmt.Println(root.Val)
}
