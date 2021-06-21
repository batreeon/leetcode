/*
 * @lc app=leetcode.cn id=1239 lang=golang
 *
 * [1239] 串联字符串的最大长度
 */

// @lc code=start
package main
import "sort"

func maxLength(arr []string) int {
	// 暴力法，找出所有的可能组合，然后比较长度找出最长的
	sort.Slice(arr,func(i,j int) bool {return len(arr[i]) > len(arr[j])})
	n := len(arr)
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	result := 0
	var tb func(letters []bool,begin,l int) 
	tb = func(letters []bool,begin,l int) {
		result = max(result,l)
		if begin == n {
			return
		}
		for i := begin ; i < n ; i++ {
			next := false
			for j := range arr[i] {
				if letters[arr[i][j]-'a'] {
					for k := j-1 ; k >= 0 ; k-- {
						letters[arr[i][k]-'a'] = false
					}
					next = true
					break
				}
				letters[arr[i][j]-'a'] = true
			}
			if next {
				continue
			}
			tb(letters,i+1,l+len(arr[i]))
			for j := range arr[i] {
				letters[arr[i][j]-'a'] = false
			}
		}
	}
	
	tb(make([]bool,26),0,0)
	return result
}
// @lc code=end

