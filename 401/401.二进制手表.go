/*
 * @lc app=leetcode.cn id=401 lang=golang
 *
 * [401] 二进制手表
 */

// @lc code=start
package main

import (
	"math"	
	"strconv"
)


func readBinaryWatch(turnedOn int) []string {
	result := []string{}
	var tb func(length int,begin int,curlen int,curval int,threslen int,thresval int,res *[]int)
	tb = func(length int,begin int,curlen int,curval int,threslen int,thresval int,res *[]int) {
		if curlen >= threslen {
			if curval <= thresval {
				*res = append(*res,curval)
			}
			return
		}
		for i := begin ; i < length ; i++ {
			tb(length,i+1,curlen+1,curval+int(math.Pow(2,float64(i))),threslen,thresval,res)
		}
	}
	for hlen := 0 ; hlen <= turnedOn ; hlen++ {
		mlen := turnedOn - hlen
		if hlen > 4 || mlen > 6 {
			continue
		}
		h,m := []int{},[]int{}
		tb(4,0,0,0,hlen,11,&h)
		tb(6,0,0,0,mlen,59,&m)
		for i := range h {
			strh := strconv.Itoa(h[i])
			for j := range m {
				strm := strconv.Itoa(m[j])
				if m[j] < 10 {
					result = append(result,strh+":0"+strm)
				}else{
					result = append(result,strh+":"+strm)
				}
			}
		}
	}
	return result
}
// @lc code=end

