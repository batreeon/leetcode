/*
 * @lc app=leetcode.cn id=1052 lang=golang
 *
 * [1052] 爱生气的书店老板
 */

// @lc code=start
package main
func maxSatisfied(customers []int, grumpy []int, X int) int {
	n := len(customers)
	left := 0
	noClam,maxClam,clam := 0,0,0
	//noClam 不抑制情绪得到的结果
	//maxClam 抑制情绪得到的最大收益
	//clam 当前窗口抑制情绪得到的收益
	for i := 0 ; i < n ; i++ {
		if grumpy[i] == 0 {
			noClam += customers[i]
		}else{
			clam += customers[i]
		}
		if i-left+1 > X {
			if grumpy[left] == 1 {
				clam -= customers[left]
			}
			left++
		}
		if clam > maxClam {
			maxClam = clam
		}
	}
	return noClam + maxClam
}
// @lc code=end

