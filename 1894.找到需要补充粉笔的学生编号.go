/*
 * @lc app=leetcode.cn id=1894 lang=golang
 *
 * [1894] 找到需要补充粉笔的学生编号
 */

// @lc code=start
package main

import (
	"sort"
)

func chalkReplacer(chalk []int, k int) int {
	/*
	n := len(chalk)
	presum := make([]int,n+1)
	for i := range chalk {
		presum[i+1] = presum[i]+chalk[i]
	}
	yu := k%presum[n]
	l,r := 1,n
	for l < r {
		mid := (l+r)/2
		if presum[mid] == yu {
			return mid
		}else if presum[mid] > yu {
			r = mid
		}else{
			l = mid+1
		}
	}
	return l-1
	*/

	n := len(chalk)
	sum := make([]int,n)
	sum[0] = chalk[0]
	for i := 1 ; i < n ; i++ {
		sum[i] = sum[i-1] + chalk[i]
	}

	k = k%sum[n-1]
	idx := sort.Search(n,func(i int) bool {return sum[i] > k})
	return idx
}
// @lc code=end

