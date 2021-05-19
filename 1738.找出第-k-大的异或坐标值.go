/*
 * @lc app=leetcode.cn id=1738 lang=golang
 *
 * [1738] 找出第 K 大的异或坐标值
 */

// @lc code=start
package main
import "sort"

func kthLargestValue(matrix [][]int, k int) int {
	// 暴力法，计算出所有结果，然后排序找到第k大的。。。
	r,c := len(matrix),len(matrix[0])
	roxs1 := make([]int,c+1)
	sortedroxs := []int{}
	for i := 1 ; i <= r ; i++ {
		roxs := make([]int,c+1)
		for j := 1 ; j <= c ; j++ {
			roxs[j] = roxs1[j-1] ^ roxs1[j] ^ roxs[j-1] ^ matrix[i-1][j-1]
			sortedroxs = append(sortedroxs,roxs[j])
		}
		roxs1 = roxs
	}
	sort.Ints(sortedroxs)
	return sortedroxs[r*c-k]
}
// @lc code=end

