/*
 * @lc app=leetcode.cn id=299 lang=golang
 *
 * [299] 猜数字游戏
 */

// @lc code=start
package main

import (
	"strconv"
	"strings")


func getHint(secret string, guess string) string {
	/*
	A,B := 0,0
	cnts := map[byte]int{}
	for i := 0 ; i < len(guess) ; i++ {
		if secret[i] == guess[i] {
			cnts[guess[i]] = strings.Count(secret,string(guess[i]))-1
			A++
		}else if strings.Contains(secret,string(guess[i])) {
			if cnt,ok := cnts[guess[i]] ; ok && cnt != 0 {
				cnts[guess[i]]--
				B++
			}else if !ok {
				cnts[guess[i]] = strings.Count(secret,string(guess[i]))-1
				B++
			}
		}
	}

	x := strconv.Itoa(A)
	y := strconv.Itoa(B)

	return strings.Join([]string{x,"A",y,"B"},"")
	*/

	A := 0
	B := 0
	cnts := map[byte]int{}
	s,g := []byte(secret),[]byte(guess)
	n := len(s)
	for i := 0 ; i < n ; {
		if s[i] == g[i] {
			copy(s[i:n-1],s[i+1:n])
			copy(g[i:n-1],g[i+1:n])
			n--
			A++
		}else{
			i++
		}
	}

	for i := 0 ; i < n ; i++ {
		if cnt,ok := cnts[g[i]] ; ok && cnt != 0 {
			cnts[g[i]]--
			B++
		}else if !ok {
			cnts[g[i]] = strings.Count(string(s[:n]),string(g[i]))
			if cnts[g[i]] != 0 {
				cnts[g[i]]--
				B++
			}
		}
	}

	x := strconv.Itoa(A)
	y := strconv.Itoa(B)

	return strings.Join([]string{x,"A",y,"B"},"")
}
// @lc code=end

