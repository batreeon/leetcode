/*
 * @lc app=leetcode.cn id=628 lang=golang
 *
 * [628] 三个数的最大乘积
 */

// @lc code=start
func maximumProduct(nums []int) int {
//结果可能有两种构造，最大的三个正数的乘积，或者最小的两个负数和最大正数的乘积
	min1,min2 := math.MaxInt64,math.MaxInt64
	max1,max2,max3 := math.MinInt64,math.MinInt64,math.MinInt64

	for _,x := range nums {
		if x < min1 {
			min2 = min1
			min1 = x
		}else if x < min2 {
			min2 = x
		}

		if x > max1 {
			max3 = max2
			max2 = max1
			max1 = x
		}else if x > max2 {
			max3 = max2
			max2 = x
		}else if x > max3 {
			max3 = x
		}
	}
	//golang 不支持三元运算符
	ans1 := min1 * min2 * max1
	ans2 := max1 * max2 * max3
	// return ans1 > ans2 ? ans1 : ans2
	return max(ans1,ans2)
}
func max(x,y int) int {
	if x > y {
		return x
	}else{
		return y
	}
}
// @lc code=end

