/*
 * @lc app=leetcode.cn id=2326 lang=golang
 *
 * [2326] 螺旋矩阵 IV
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func spiralMatrix(m int, n int, head *ListNode) [][]int {
    result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = -1
		}
	}

	dirs := [][]int{{0,1}, {1,0}, {0, -1}, {-1, 0}}
	di := 0
	i, j := 0, 0
	for head != nil {
		result[i][j] = head.Val
		nexti, nextj := i + dirs[di][0], j + dirs[di][1]
		if nexti < 0 || nexti >= m || nextj < 0 || nextj >= n || result[nexti][nextj] != -1 {
			di = (di + 1) % 4
		}
		i, j = i + dirs[di][0], j + dirs[di][1]
		head = head.Next
	}
	return result
}
// @lc code=end

