/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
package main
func topKFrequent(nums []int, k int) []int {
	//记录每个数的重复次数,key是值，value是次数
	m := map[int]int{}
	for _,n := range nums {
		m[n]++
	}
	
	times := [][]int{}
	for k,v := range m {
		times = append(times,[]int{k,v})
	}

	partition := func(times [][]int,l,r int) int {
		x := times[r][1]
		i := l - 1
		for j := l ; j < r ; j++ {
			if times[j][1] >= x {//大的排前面
				i++
				times[i],times[j] = times[j],times[i]
			}
		}
		times[i+1],times[r] = times[r],times[i+1]
		return i+1
	}

	start,end := 0,len(times)-1
	p := partition(times,start,end)
	for p != k-1 {
		if p > k-1 {
			end = p-1
		}else{
			start = p+1
		}
		p = partition(times,start,end)
	}
	ans := []int{}
	for i := 0 ; i < k ; i++ {
		ans = append(ans,times[i][0])
	}
	return ans
}
// @lc code=end

