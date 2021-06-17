/*
 * @lc app=leetcode.cn id=65 lang=golang
 *
 * [65] 有效数字
 */

// @lc code=start
package main
// import "strings"


func isNumber(s string) bool {
	/*
	ne := strings.Count(s,"e") + strings.Count(s,"E")
	if ne > 1 {
		return false
	}
	ndot := strings.Count(s,".")
	if ndot > 1 {
		return false
	}
	isnum := func(ss string) bool {
		if len(ss) == 0 {
			return false
		}
		nsign := strings.Count(ss,"+") + strings.Count(ss,"-")
		if nsign > 1 {
			return false
		}
		if nsign == 1 {
			if len(ss) == 1 {
				return false
			}
			if ss[0] != '+' && ss[0] != '-' {
				return false
			}
			ss = ss[1:]
		}
		ndot := strings.Count(ss,".")
		if ndot > 1 {
			return false
		}else if ndot == 1 {
			if len(ss) == 1 {
				return false
			}
		}
		return true
	}
	if ne == 1 {
		ie := strings.IndexAny(s,"eE")
		if ie == len(s)-1 {
			return false
		}
		l,r := s[:ie],s[ie+1:]
		if strings.Contains(r,".") {
			return false
		}
		return isnum(l) && isnum(r)
	}else{
		return isnum(s)
	}
	return false
	*/

	n := len(s)
	ie := -1
	for i := 0 ; i < n ; i++ {
		if s[i] == 'e' || s[i] == 'E' {
			if ie == -1 {
				ie = i
			}else{
				return false
			}
		}
	}

	var check func(ss string,mustinteger bool) bool
	check = func(ss string,mustinteger bool) bool {
		if len(ss) == 0 {
			return false
		}
		if ss[0] == '+' || ss[0] == '-' {
			ss = ss[1:]
		}
		hasDot,hasNum := false,false
		for i := 0 ; i < len(ss) ; i++ {
			if ss[i] == '.' {
				if hasDot || mustinteger {
					return false
				}
				hasDot = true
			}else if ss[i] >= '0' && ss[i] <= '9' {
				hasNum = true
			}else{
				return false
			}
		}
		return hasNum
	}

	result := true
	if ie != -1 {
		result = result && check(s[:ie],false)
		result = result && check(s[ie+1:],true)
	}else{
		result = result && check(s,false)
	}
	return result
}
// @lc code=end

