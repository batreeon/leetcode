/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	// 只有这一种溢出可能，考虑除数，被除数，商都应为有符号32位整数
	if dividend == (-1)*(2<<30) && divisor == -1 {
		return 2<<30 - 1
	}
	// 例子60/8=
	// 8*2*2 <= 60,8*2*2*2 > 60
	// 60 - 32 = 28;28 / 8=
	// 8*2 <= 28 ; 8*2*2 > 28
	// 28 - 16 = 12; 12 / 8 = 
	// 8*2 > 12
	// 则最终结果= 4+2+1 = 7
	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}
	if dividend < 0 {
		dividend = -1 * dividend
	}
	if divisor < 0 {
		divisor = -1 * divisor
	}
	ans := 0
	for dividend >= divisor {
		res := divisor
		exponent := 1	//指数
		for res << 1 <= dividend {//提前判断，找出上界
			res <<= 1
			exponent <<= 1	//指数是成指数增长的
		}
		ans += exponent
		dividend -= res
	}
	return ans*sign
}
// @lc code=end

