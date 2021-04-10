/*
 * @lc app=leetcode.cn id=1361 lang=golang
 *
 * [1361] 验证二叉树
 */

// @lc code=start
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	in := make([]int,n)
	color := make([]int,n)
	colorNum := 0

	for i := 0 ; i < n ; i++ {
		if color[i] == 0 {//未标色
			colorNum++
			color[i] = colorNum
		}

		leftchild,rightchild := leftChild[i],rightChild[i]
		
		if leftchild != -1 {
			in[leftchild]++
			if in[leftchild] > 1 {
				return false
			}
			
			if color[leftchild] == 0 {
				color[leftchild] = color[i]
			}else if color[leftchild] == color[i] {
				// 颜色相同说明有环
				return false
			}
		}
		if rightchild != -1 {
			in[rightchild]++
			if in[rightchild] > 1 {
				return false
			}

			if color[rightchild] == 0 {
				color[rightchild] = color[i]
			}else if color[rightchild] == color[i] {
				// 颜色相同说明有环
				return false
			}
		}
	}

	sum := 0
	for i := 0 ; i < n ; i++ {
		sum += in[i]
	}
	// 防止有多棵树，检查入度和
	return sum == n-1
}
// @lc code=end

