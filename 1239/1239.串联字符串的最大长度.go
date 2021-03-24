 /*
 * @lc app=leetcode.cn id=1239 lang=golang
 *
 * [1239] 串联字符串的最大长度
 */

// @lc code=start
package main

import "sort"

type sortedString []string

func (s sortedString) Len() int {
	return len(s)
}

// 按长度递减排序
func (s sortedString) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}
func (s sortedString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func maxLength(arr []string) int {
	// 应该用递归的
	n := len(arr)
	if n <= 1 {
		return len(arr[0])
	}
	sortedArr := sortedString(arr)
	sort.Sort(sortedArr)
	maxAns := 0
	ans := 0
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	seen := map[rune]bool{}
	var backTrack func(arr []string)
	backTrack = func(arr []string) {
		if len(arr) == 0 {
			return
		}
		for i := 0 ; i < len(arr) ; i++ {
			// 一共有26个字母，总字母数比这个长，肯定是不符合的
			if len(arr[i])+ans > 26 {
				continue
			}
			isContinue := false
			for j,k := range arr[i] {
				if seen[k] == true {
					isContinue = true
					// 若在当前下标j时，字母重复，
					// 那么arr[i]中j之前的元素都是只出现一次，需要从map中去掉
					j--
					for j >= 0 {
						delete(seen,rune(arr[i][j]))
						j--
					}
					break
				}
				seen[k] = true
			}
			if isContinue {
				continue
			}
			ans += len(arr[i])
			maxAns = max(maxAns,ans)
			backTrack(arr[i+1:])
			for _,k := range arr[i] {
				delete(seen,k)
			}
			ans -= len(arr[i])
		}
	}
	backTrack(sortedArr)
	return maxAns
}

// @lc code=end
