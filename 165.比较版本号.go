/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] 比较版本号
 */

// @lc code=start
package main

import (
	"strconv"
	"strings"
)

/*
func compareVersion(version1 string, version2 string) int {
	f := func(r rune)bool{
		if r == rune('.') {
			return true
		}
		return false
	}
	v1,v2 := strings.FieldsFunc(version1,f),strings.FieldsFunc(version2,f)
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	l1,l2 := len(v1),len(v2)
	maxIndex := min(l1,l2)
	isZero := func(r rune)bool {
		if r == rune('0') {
			return true
		}
		return false
	}
	for i := 0 ; i < maxIndex ; i++ {
		num1s,num2s := strings.TrimLeftFunc(v1[i],isZero),strings.TrimLeftFunc(v2[i],isZero)
		num1,_ := strconv.Atoi(num1s)
		num2,_ := strconv.Atoi(num2s)
		if num1 > num2 {
			return 1
		}else if num2 > num1 {
			return -1
		}
	}
	if l1 == l2 {
		return 0
	}
	if l1 > l2 {
		for i := maxIndex ; i < l1 ; i++ {
			if strings.TrimLeftFunc(v1[i],isZero) != "" {
				return 1
			}
		}
	}else{
		for i := maxIndex ; i < l2 ; i++ {
			if strings.TrimLeftFunc(v2[i],isZero) != "" {
				return -1
			}
		}
	}
	return 0
}
*/
func compareVersion(version1 string, version2 string) int {
	v1s,v2s := strings.Split(version1,"."),strings.Split(version2,".")
	maxIndex := len(v1s)
	if len(v2s) < maxIndex {
		maxIndex = len(v2s)
	}
	for i := 0 ; i < maxIndex ; i++ {
		num1,_ := strconv.Atoi(strings.TrimLeft(v1s[i],"0"))
		num2,_ := strconv.Atoi(strings.TrimLeft(v2s[i],"0"))
		if num1 < num2 {
			return -1
		}else if num1 > num2 {
			return 1
		}
	}
	if len(v1s) == len(v2s) {
		return 0
	}
	if len(v1s) < len(v2s) {
		for i := maxIndex ; i < len(v2s) ; i++ {
			if strings.TrimLeft(v2s[i],"0") != "" {
				return -1
			}
		}
	}else{
		for i := maxIndex ; i < len(v1s) ; i++ {
			if strings.TrimLeft(v1s[i],"0") != "" {
				return 1
			}
		}
	}
	return 0
}
// @lc code=end

