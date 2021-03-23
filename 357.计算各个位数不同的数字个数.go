/*
 * @lc app=leetcode.cn id=357 lang=golang
 *
 * [357] 计算各个位数不同的数字个数
 */

// @lc code=start
package main

func countNumbersWithUniqueDigits(n int) int {
	/*
	// 多放了几个现成解
	ans := []int{1,10,91,739,5275}
	if n <= 4 {
		return ans[n]
	}
	for i := 5 ; i <= n ; i++ {
		temp := 1
		for j := 9 ; j > 9-i+1 ; j-- {
			temp *= j
		}
		ans = append(ans,ans[i-1]+temp*9)
	}
	return ans[n]
	*/

	// 只存第一个，但结果成绩没差别
	ans := 1
	if n > 0 {
		for i := 1 ; i <= n ; i++ {
			temp := 1
			for j := 9 ; j > 9-i+1 ; j-- {
				temp *= j
			}
			ans += temp*9
		}
	}
	return ans
}
// @lc code=end

