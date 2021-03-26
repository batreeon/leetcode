/*
 * @lc app=leetcode.cn id=556 lang=golang
 *
 * [556] 下一个更大元素 III
 */

// @lc code=start
package main
import "sort"

func nextGreaterElement(n int) int {
	nn := []int{}
	for n != 0 {
		nn = append(nn,n%10)
		n /= 10
	}
	l := len(nn)
	for i,j := 0,l-1 ; i < j ; i,j = i+1,j-1 {
		nn[i],nn[j] = nn[j],nn[i]
	}
	i := l-1
	for  ; i > 0 ; i-- {
		if nn[i-1] < nn[i] {
			break
		}
	}
	if i == 0 {
		return -1
	}

	// index记录
	index := i
	
	for j := i+1 ; j < l ; j++ {
		if nn[j] > nn[i-1] {
			index = j
		}else{
			break
		}
	}
	nn[i-1],nn[index] = nn[index],nn[i-1]

	copyN := make([]int,l-i)
	copy(copyN,nn[i:])
	sort.Ints(copyN)
	copy(nn[i:],copyN)
	ans := 0
	for _,num := range nn {
		ans = 10*ans+num
	}
	if ans > (1<<31)-1 {
		return -1
	}
	return ans
}
// @lc code=end

