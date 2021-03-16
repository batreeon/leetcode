/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}//没有这句，如果走循环的话，""的输出是[""],正确结果应为[]
	letters := make([]string,10)
	letters[2] = "abc"
	letters[3] = "def"
	letters[4] = "ghi"
	letters[5] = "jkl"
	letters[6] = "mno"
	letters[7] = "pqrs"
	letters[8] = "tuv"
	letters[9] = "wxyz"
	ans := []string{}
	var lc func(digits string,res []byte) 
	lc = func(digits string,res []byte) {
		if len(digits) == 0 {
			ans = append(ans,string(res))
		}else{
			index := digits[0] - '0'
			for _,l := range letters[index] {
				res = append(res,byte(l))
				lc(digits[1:],res)
				res = res[:len(res)-1]
			}
		}
	}
	lc(digits,[]byte{})
	return ans
}
// @lc code=end

