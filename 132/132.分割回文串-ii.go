/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 */

// @lc code=start
package main
import "math"

func minCut(s string) int {
	n := len(s)
	// f用来记录所有子串的情况，所有i,j对是否为回文串
	f := make([][]bool,n)
	fRow := make([]bool,n)
	for i := 0 ; i < n ; i++ {
		fRow[i] = true
	}
	for i := 0 ; i < n ; i++ {
		f[i] = make([]bool,n)
		copy(f[i],fRow)
	}
	for i := n-1 ; i >= 0 ; i-- {
		for j := i+1 ; j < n ; j++ {
			f[i][j] = ((s[i] == s[j]) && f[i+1][j-1])
		}
	}
	/*
	times := 0
	minTimes := math.MaxInt64
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	var minCutPalindrome func(l,r int,times *int)
	minCutPalindrome = func(l,r int,times *int) {
		if f[l][r] {
			minTimes = min(minTimes,*times)
			return
		}
		for i := r - 1 ; i >= l ; i-- {
			*times++
			if *times >= minTimes {
				*times--
				return
			}
			if f[l][i] {
				minCutPalindrome(i+1,r,times)
			}
			*times--
		}
	}
	minCutPalindrome(0,n-1,&times)
	return minTimes
	*/
	//不能再用递归了，现在用例很长
	// 从小到大计算g[i],0<=i<=n-1,表示[0,i]序列的最小切割次数
	
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	g := make([]int,n)
	for i := 0 ; i < n ; i++ {
		g[i] = math.MaxInt64
		if f[0][i] {
			g[i] = 0
		}else{
			for j := 0 ; j < i ; j++ {
				if f[j+1][i] {
					g[i] = min(g[i],g[j]+1)
				}
			}
		}
	}
	return g[n-1]
}
// @lc code=end

