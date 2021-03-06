/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
package main

func isUgly(num int) bool {
	/*
	var isugly func(x int) bool 
	isugly = func(x int) bool {
		if x <= 0 {
			return false
		}//丑数是正数
		
		if x <= 5 {
			return true
		}else{
			if x % 2 == 0 {
				return isugly(x/2)
			}else if x % 3 == 0 {
				return isugly(x/3)
			}else if x % 5 == 0 {
				return isugly(x/5)
			}else{
				return false
			}
		}
		
	}
	return isugly(num)
	*/
	x := num
	if x <= 0 {
		return false
	}
	for x % 2 == 0 {
		x /= 2
	}
	for x % 3 == 0 {
		x /= 3
	}
	for x % 5 == 0 {
		x /= 5
	}
	return x == 1
}
// @lc code=end

