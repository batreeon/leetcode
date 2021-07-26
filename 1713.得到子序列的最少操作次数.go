/*
 * @lc app=leetcode.cn id=1713 lang=golang
 *
 * [1713] 得到子序列的最少操作次数
 */

// @lc code=start
package main
import "sort"
func minOperations(target []int, arr []int) int {
	/*
	n1,n2 := len(target),len(arr)
	f := make([]int,n2+1)
	for i := 0 ; i < n1 ; i++ {
		f1 := make([]int,n2+1)
		for j := 0 ; j < n2 ; j++ {
			if target[i] == arr[j] {
				f1[j+1] = f[j]+1
			}else{
				if f1[j] > f[j+1] {
					f1[j+1] = f1[j]
				}else{
					f1[j+1] = f[j+1]
				}
			}
		}
		f = f1
	}
	return n1-f[n2]
	*/
	pos := map[int]int{}
	for i := range target {
		pos[target[i]] = i
	}
	d := []int{}
	for i := range arr {
		if idx,ok := pos[arr[i]] ; ok {
			id := sort.Search(len(d),func(j int) bool {return d[j] >= idx})
			if id == len(d) {
				d = append(d,idx)
			}else{
				d[id] = idx
			}
		}
	}
	return len(target) - len(d)
}
// @lc code=end

