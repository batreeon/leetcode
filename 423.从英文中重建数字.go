/*
 * @lc app=leetcode.cn id=423 lang=golang
 *
 * [423] 从英文中重建数字
 */

// @lc code=start
package main

/*
// 用于参考，这个表没在程序里用
var m map[string]int = map[string]int{
	"zero": 0,
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
}
*/

func originalDigits(s string) string {
	letters := map[int]int{}
	for _, letter := range s {
		letters[int(letter-'a')]++
	}
	result := make([]int, 10)

	if n := letters[int('z'-'a')]; n > 0 {
		result[0] = n
		for _,b := range "zero" {
			letters[int(b-'a')] -= n
		}
	}
	if n := letters[int('w'-'a')]; n > 0 {
		result[2] = n
		for _,b := range "two" {
			letters[int(b-'a')] -= n
		}
	}
	if n := letters[int('u'-'a')]; n > 0 {
		result[4] = n
		for _,b := range "four" {
			letters[int(b-'a')] -= n
		}
	}
	if n := letters[int('f'-'a')]; n > 0 {
		result[5] = n
		for _,b := range "five" {
			letters[int(b-'a')] -= n
		}
	}
	if n := letters[int('x'-'a')]; n > 0 {
		result[6] = n
		for _,b := range "six" {
			letters[int(b-'a')] -= n
		}
	}
	if n := letters[int('v'-'a')]; n > 0 {
		result[7] = n
		for _,b := range "seven" {
			letters[int(b-'a')] -= n
		}
	}
	if n := letters[int('g'-'a')]; n > 0 {
		result[8] = n
		for _,b := range "eight" {
			letters[int(b-'a')] -= n
		}
	}
	if n := letters[int('o'-'a')]; n > 0 {
		result[1] = n
		for _,b := range "one" {
			letters[int(b-'a')] -= n
		}
	}
	if n := letters[int('h'-'a')]; n > 0 {
		result[3] = n
		for _,b := range "two" {
			letters[int(b-'a')] -= n
		}
	}
	result[9] = letters[int('i'-'a')]

	res := string("")
	for i,n := range result {
		for n > 0 {
			res += string('0'+i)
			n--
		}
	}

	return res
}
// @lc code=end

