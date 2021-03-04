/*
 * @lc app=leetcode.cn id=331 lang=golang
 *
 * [331] 验证二叉树的前序序列化
 */

// @lc code=start
package main
import "strings"
func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder,",")
	if len(nodes) == 0 || (len(nodes) == 1 && nodes[0] == "#") {
		return true
	}
	if len(nodes) > 1 && nodes[0] == "#" {
		return false
	}

	stack := []string{nodes[0]}
	is := 0
	delete := []string{}
	id := -1
	var deleteNode func()

	deleteNode = func() {
		stack = stack[:is]
		is--
		delete = delete[:id]
		id--
		if id >= 0 && delete[id] == stack[is] {
			deleteNode()
		}else{
			// 删除的时候，还要检查下一步，若上一步删除的是一个右儿子，那么还要继续删掉前面的节点，
			// 若上一步删除的是一个左儿子，那么就把上一个节点加入待删栈
			if is >= 0 {
				delete = append(delete,stack[is])
				id++
			}
		}
	}

	for i := 1 ; i < len(nodes) ; i++ {
		if is < 0 {
			return false
		}
		if nodes[i] != "#" {
			stack = append(stack,nodes[i])
			is++
		}
		if nodes[i] == "#" {
			if id >= 0 && delete[id] == stack[is] {
				deleteNode()
			}else{
				delete = append(delete,stack[is])
				id++
			}
		}
	}
	return is < 0 
}
// @lc code=end

