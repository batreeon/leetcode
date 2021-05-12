/*
 * @lc app=leetcode.cn id=1734 lang=golang
 *
 * [1734] 解码异或后的排列
 */

// @lc code=start
func decode(encoded []int) []int {
	total := 0
	n := len(encoded)+1
	for i := 1 ; i <= n ; i++ {
		total ^= i
	}
	
	// suf是剔除perm[0]后其他所有的元素的xor结果
	suf := 0
	for i := 1 ; i < n-1 ; i += 2 {
		suf ^= encoded[i]
	}
	perm := make([]int,n)
	perm[0] = total ^ suf

	for i := 1 ; i < n ; i++ {
		perm[i] = encoded[i-1] ^ perm[i-1]
	}
	return perm
}
// @lc code=end

