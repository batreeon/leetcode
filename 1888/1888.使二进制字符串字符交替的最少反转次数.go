/*
 * @lc app=leetcode.cn id=1888 lang=golang
 *
 * [1888] 使二进制字符串字符交替的最少反转次数
 */

// @lc code=start
package main
func minFlips(s string) int {
	/*
	// 错了，"01001001101"没通过
	if len(s) < 2 {
		return 0
	}
	diffi := 0
	for i := 1 ; i < len(s) ; i++ {
		if s[i] != s[0] {
			diffi = i
			break
		}
	}
	if diffi == 0 {
		return len(s)/2
	}
	ss := []byte(s)
	temp := make([]byte,len(ss))
	copy(temp[:len(ss)-(diffi-1)],ss[diffi-1:])
	copy(temp[len(ss)-(diffi-1):],ss[:diffi-1])
	result := 0
	for i := 1 ; i < len(temp) ; i++ {
		if temp[i] == temp[i-1] {
			temp[i] = temp[i]^1
			result++
		}
	}
	return result
	*/

	n := len(s)
	// pre记录[0,i]序列以0或1结尾时需要反转的次数
	// suf记录[i,n-1]序列以0或1起始时需要反转的次数
	pre := make([][]int,n)
	I := func(x,y byte) int {
		if x == y {
			return 0
		}
		return 1
	}
	pre[0] = make([]int,2)
	pre[0][0] = I(s[0],'0')
	pre[0][1] = I(s[0],'1')
	for i := 1 ; i < n ; i++ {
		pre[i] = make([]int,2)
		pre[i][0] = pre[i-1][1]+I(s[i],'0')
		pre[i][1] = pre[i-1][0]+I(s[i],'1')
	}
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	res := min(pre[n-1][0],pre[n-1][1])
	if n&1 == 1 {//序列长度为奇数
		suf := make([][]int,n)
		suf[n-1] = make([]int,2)
		suf[n-1][0] = I(s[n-1],'0')
		suf[n-1][1] = I(s[n-1],'1')
		for i := n-2 ; i >= 0 ; i-- {
			suf[i] = make([]int,2)
			suf[i][0] = suf[i+1][1] + I(s[i],'0')
			suf[i][1] = suf[i+1][0] + I(s[i],'1')
		}
		for i := 0 ; i < n-1 ; i++ {
			res = min(res,pre[i][0]+suf[i+1][0])
			res = min(res,pre[i][1]+suf[i+1][1])
		}
	}
	return res
}
// @lc code=end

