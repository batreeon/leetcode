/*
 * @lc app=leetcode.cn id=1449 lang=golang
 *
 * [1449] 数位成本和为目标值的最大数字
 */

// @lc code=start
package main

import (
	// "sort"
	// "math"
	"strconv"
)



func largestNumber(cost []int, target int) string {
	/*
	m := make(map[int]int)
	costs := []int{}
	for i,c := range cost {
		if _,ok := m[c] ; !ok {
			costs = append(costs,c)
			m[c] = i+1
		}else{
			if i+1 > m[c] {
				m[c] = i+1
			}
		}
	}
	sort.Ints(costs)

	costss := [][]int{}
	maxlen := 0
	var tb func(int,[]int,int)
	tb = func(beg int,res []int,totalcost int) {
		if totalcost == target {
			if len(res) < maxlen {
				return
			}
			if len(res) > maxlen {
				maxlen = len(res)
				costss = costss[:0]
			}
			newres := make([]int,len(res))
			copy(newres,res)
			costss = append(costss,newres)
			return
		}
		if totalcost > target {
			return
		}
		for i := beg ; i < len(costs) ; i++ {
			res = append(res,costs[i])
			tb(i,res,totalcost+costs[i])
			res = res[:len(res)-1]
		}
	}

	tb(0,[]int{},0)

	nums := [][]int{}
	for i := range costss {
		res := []int{}
		for j := range costss[i] {
			res = append(res,m[costss[i][j]])
		}
		sort.Slice(res,func(i,j int) bool {return res[i]>res[j]})
		nums = append(nums,res)
	}

	maxV := string("")
	max := func(x,y string)string {
		if x > y {
			return x
		}
		return y
	}
	for _,num := range nums {
		res := string("")
		for _,n := range num {
			ns := strconv.Itoa(n)
			res += ns
		}
		maxV = max(maxV,res)
	}

	if maxV == string("") {
		return string("0")
	}
	return maxV
	*/

	f := make([]int,target+1)
	for i := range f {
		f[i] = -1
	}
	f[0] = 0
	for i := 1 ; i <= 9 ; i++ {
		c := cost[i-1]
		for j := c ; j <= target ; j++ {
			if f[j-c]+1 > f[j] && f[j-c]+1 > 0 {
				f[j] = f[j-c]+1
			}
		}
	}
	if f[target] <= 0 {
		return string("0")
	}
	result := string("")
	for i,j := 9,target ; i >= 1 ; i-- {
		c := cost[i-1]
		for j >= c && f[j] == f[j-c]+1 {
			// f[j] == f[j-c]+1这个判断好关键啊
			result += strconv.Itoa(i)
			j -= c
		}
	}
	return result
}
// @lc code=end

