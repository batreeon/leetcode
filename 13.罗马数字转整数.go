/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	ss := []byte(s)
	m := map[string]int{}
	m["I"] = 1
	m["IV"] = 4
	m["V"] = 5
	m["IX"] = 9
	m["X"] = 10
	m["XL"] = 40
	m["L"] = 50
	m["XC"] = 90
	m["C"] = 100
	m["CD"] = 400
	m["D"] = 500
	m["CM"] = 900
	m["M"] = 1000
	result := 0
	for len(ss) > 0 {
		if len(ss) >= 2 {
			if v,ok := m[string(ss[len(ss)-2:])] ; ok {
				result += v
				ss = ss[:len(ss)-2]
			}else{
				result += m[string(ss[len(ss)-1:])]
				ss = ss[:len(ss)-1]
			}
		}else{
			result += m[string(ss[:])]
			return result
		}
	}
	return result
}
// @lc code=end

