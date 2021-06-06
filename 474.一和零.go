/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start
package main
import "strings"


func findMaxForm(strs []string, m int, n int) int {
	/*
	length := len(strs)
	cnt := make([][]int,length)
	for i := range cnt {
		cnt[i] = make([]int,2)
	}
	for i := range strs {
		cnt[i][0],cnt[i][1] = strings.Count(strs[i],"0"),strings.Count(strs[i],"1")
	}

	f := make([][]int,m+1)
	for i := range f {
		f[i] = make([]int,n+1)
	}
	for k := 0 ; k < length ; k++ {
		zero,one := cnt[k][0],cnt[k][1]
		for i := m ; i >= zero ; i-- {
			for j := n ; j >= one ; j-- {
				f[i][j] = max(f[i][j],f[i-zero][j-one]+1)
			}
		}
	}
	return f[m][n]
	*/

	length := len(strs)
	cnt := make([][]int,length)
	for i := range cnt {
		cnt[i] = make([]int,2)
	}
	for i := range strs {
		cnt[i][0],cnt[i][1] = strings.Count(strs[i],"0"),strings.Count(strs[i],"1")
	}

	f := make([][]int,m+1)
	for i := range f {
		f[i] = make([]int,n+1)
	}
	for i := 0 ; i <= m ; i++ {
		for j := 0 ; j <= n ; j++ {
			if i >= cnt[0][0] && j >= cnt[0][1] {
				f[i][j] = 1
			}
		}
	}
	for k := 1 ; k < length ; k++ {
		zero,one := cnt[k][0],cnt[k][1]
		f1 := make([][]int,m+1)
		for i := range f1 {
			f1[i] = make([]int,n+1)
		}
		for i := 0 ; i <= m ; i++ {
			for j := 0 ; j <= n ; j++ {
				a := f[i][j]
				b := 0
				if i >= zero && j >= one {
					b = f[i-zero][j-one]+1
				}
				f1[i][j] = max(a,b)
			}
		}
		f = f1
	}
	return f[m][n]
}
func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

