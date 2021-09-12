/*
 * @lc app=leetcode.cn id=600 lang=golang
 *
 * [600] 不含连续1的非负整数
 */

// @lc code=start
package main

func findIntegers(n int) int {
	/*
	// 不会做
	if n < 3 {
		return n+1
	}
	length := 1
	for (n >> length) > 0 {
		length++
	}
	dp := make([]int,length+1)
	dp[1] = 2
	for i := 2 ; i <= length ; i++ {
		dp[i] = (1 << i) - ((1 << (i-1))-dp[i-1]) - (1 << (i-2)) 
	}
	result := dp[length-1]
	result += n - (1 << (length-1)) + 1
	for n >= (1 << (length-1)) {
		nn := n
		for nn >= 3 {
			if nn & 3 == 3 {
				result--
				break
			}
			nn >>= 1
		}
		n -= 1
	}
	return result
	*/

	// 抄的
	ans := 0
	dp := [31]int{1, 1}
    for i := 2; i < 31; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    for i, pre := 29, 0; i >= 0; i-- {
        val := 1 << i
        if n&val > 0 {
            ans += dp[i+1]
            if pre == 1 {
                break
            }
            pre = 1
        } else {
            pre = 0
        }
        if i == 0 {
            ans++
        }
    }
    return ans
}

// @lc code=end

