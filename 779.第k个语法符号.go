/*
 * @lc app=leetcode.cn id=779 lang=golang
 *
 * [779] 第K个语法符号
 */

// @lc code=start
package main
func kthGrammar(N int, K int) int {
	var kthgrammar func(K int) int
	kthgrammar = func(K int) int {
		// 其实前n位都是固定的，和层数无关，
		// 你甚至可以建一个很大的表，将n位的结果都存起来
		switch K {
		case 0:
			fallthrough
		case 3:
			fallthrough
		case 5:
			fallthrough
		case 6:
			return 0
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 4:
			fallthrough
		case 7:
			return 1
		}

		// 每一层2n,2n+1位（下标从0开始记），取决于上一层第n位
		if kthgrammar(K/2) == 0 {
			if K&1 == 0 {
				return 0
			}else{
				return 1
			}
		}else{
			if K&1 == 0 {
				return 1
			}else{
				return 0
			}
		}
	}
	return kthgrammar(K-1)
}
// @lc code=end

