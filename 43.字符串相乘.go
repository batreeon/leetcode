/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	//仿竖式
	//来源：https://leetcode-cn.com/problems/multiply-strings/solution/you-hua-ban-shu-shi-da-bai-994-by-breezean/
	if num1 == string("0") || num2 == string("0") {
		return string("0")
	}
	n1,n2 := len(num1),len(num2)
	res := make([]int,n1+n2)//高位在左，小下标
	for i := n1-1 ; i >= 0 ; i-- {
		v1 := num1[i] - '0'
		for j := n2 - 1 ; j >= 0 ; j-- {
			v2 := num2[j] - '0'
			sum := res[i+j+1]+int(v1*v2)
			res[i+j+1] = sum%10
			res[i+j] += sum/10
		}
	}
	resByte := []byte{}
	for i,v := range res {
		if i == 0 && v == 0 {
			continue
		}
		resByte = append(resByte,byte('0'+v))
	}
	return string(resByte)
}
// @lc code=end

