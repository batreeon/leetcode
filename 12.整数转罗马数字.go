/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
func intToRoman(num int) string {
	norman := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
	nums := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}

	var result string
	for i := 0 ; i < 13 ; i++ {
		for num >= nums[i] {
			num -= nums[i]
			result += norman[i]
		}
	}
	return result
}
// @lc code=end

