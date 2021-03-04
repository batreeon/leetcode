/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
package main

import (
	"strconv"
)

func decodeString(s string) string {
	nums := []int{}
	brackets := []byte{}
	codes := [][]byte{}
	ans := []byte{}
	l := len(s)
	for i := 0 ; i < l ; {
		if s[i] - 'a' < 26 {
			res := []byte{s[i]}
			i++
			for ; i < l && s[i] - 'a' < 26 ; i++ {
				res = append(res,s[i])
			}
			if len(nums) == 0 {
				ans = append(ans,res...)
			}else{
				codes[len(codes)-1] = append(codes[len(codes)-1],res...)
			}
		}else if s[i] - '0' < 10 {
			res := []byte{s[i]}
			i++
			for ; s[i] - '0' < 10 ; i++ {
				res = append(res,s[i])
			}
			num,_ := strconv.Atoi(string(res))
			nums = append(nums,num)
		}else if s[i] == '[' {
			brackets = append(brackets, '[')
			i++
			res := []byte{}
			for ; s[i] - 'a' < 26 ; i++ {
				res = append(res,s[i])
			}
			codes = append(codes,res)
		}else if s[i] == ']' {
			nl := len(nums)
			num := nums[nl-1]
			nums = nums[:nl-1]

			nb := len(brackets)
			brackets = brackets[:nb-1]

			nc := len(codes)
			lc := len(codes[nc-1])
			for i := 1 ; i < num ; i++ {
				codes[nc-1] = append(codes[nc-1],codes[nc-1][:lc]...)
			}
			if nb > 1 && brackets[nb-2] == '[' {
				codes[nc-2] = append(codes[nc-2],codes[nc-1]...)
			}else{
				ans = append(ans,codes[nc-1]...)
			}
			codes = codes[:nc-1]
			i++
		}
	}
	return string(ans)
}
// @lc code=end

