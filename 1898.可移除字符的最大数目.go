/*
 * @lc app=leetcode.cn id=1898 lang=golang
 *
 * [1898] 可移除字符的最大数目
 */

// @lc code=start
package main
import "sort"


func maximumRemovals(s string, p string, removable []int) int {
	// 找到第一个使p不是s子序列的k值
	// 也就是取removable[:k+1]前k+1项时，p恰好不是s的子序列
	// 取removable[:k]前k项时，p还是s的子序列
	// 二分查找
	return sort.Search(len(removable),func(k int) bool {
		del := make([]bool,len(s))
		// 删掉前k个
		for _,i := range removable[:k+1] {
			del[i] = true
		}
		j := 0
		for i := range s {
			if !del[i] && s[i] == p[j] {
				if j++ ; j == len(p) {
					// 还是子序列
					return false
				}
			}
		}
		return true
	})
}
// @lc code=end

