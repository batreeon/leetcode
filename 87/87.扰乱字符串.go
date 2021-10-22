/*
 * @lc app=leetcode.cn id=87 lang=golang
 *
 * [87] 扰乱字符串
 */

// @lc code=start
package main
func isScramble(s1 string, s2 string) bool {
	issame := func(m1,m2 map[int]int) bool {
		for k1,v1 := range m1 {
			if v1 != m2[k1] {
				return false
			}
		}
		return true
	}
	var isscramble func(s1,s2 string,letters1,letters2 map[int]int) bool
	isscramble = func(s1,s2 string,letters1,letters2 map[int]int) bool {
		if len(s1) == 1 && len(s2) == 1 {
			return s1 == s2 
		}

		left1 := make(map[int]int)
		right1 := make(map[int]int)
		for k1,v1 := range letters1 {
			right1[k1] = v1
		}
		left2 := make(map[int]int)
		left22 := make(map[int]int)
		right2 := make(map[int]int)
		right22 := make(map[int]int)
		for k2,v2 := range letters2 {
			right2[k2] = v2
			left22[k2] = v2
		}

		for i := 0 ; i < len(s1)-1 ; i++ {
			left1[int(s1[i]-'a')]++
			left2[int(s2[i]-'a')]++

			right1[int(s1[i]-'a')]--
			if right1[int(s1[i]-'a')] == 0 {
				delete(right1,int(s1[i]-'a'))
			}
			right2[int(s2[i]-'a')]--
			if right2[int(s2[i]-'a')] == 0 {
				delete(right2,int(s2[i]-'a'))
			}

			if len(left1) == len(left2) && len(right1) == len(right2) {
				if issame(left1,left2) && issame(right1,right2) {
					if isscramble(s1[:i+1],s2[:i+1],left1,left2) && isscramble(s1[i+1:],s2[i+1:],right1,right2) {
						return true
					}
				}
			}
			// if len(left1) == len(right2) && len(right1) == len(left2) {
			// 	if issame(left1,right2) && issame(right1,left2) {
			// 		if isscramble(s1[:i+1],s2[i+1:],left1,right2) && isscramble(s1[i+1:],s2[:i+1],right1,left2) {
			// 			return true
			// 		}
			// 	}
			// }

			right22[int(s2[len(s2)-1-i]-'a')]++
			left22[int(s2[len(s2)-1-i]-'a')]--
			if left22[int(s2[len(s2)-1-i]-'a')] == 0 {
				delete(left22,int(s2[len(s2)-1-i]-'a'))
			}
			if len(left1) == len(right22) && len(right1) == len(left22) {
				if issame(left1,right22) && issame(right1,left22) {
					if isscramble(s1[:i+1],s2[len(s2)-1-i:],left1,right22) && isscramble(s1[i+1:],s2[:len(s2)-1-i],right1,left22) {
						return true
					}
				}
			}
		}
		return false
	}

	letters1,letters2 := make(map[int]int),make(map[int]int)
	for i := 0 ; i < len(s1) ; i++ {
		letters1[int(s1[i]-'a')]++
		letters2[int(s2[i]-'a')]++
	}
	return isscramble(s1,s2,letters1,letters2)
}
// @lc code=end

