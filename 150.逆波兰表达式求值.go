/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
	stack := []int{}
	isNum := func(s string) bool {
		_, err := strconv.ParseFloat(s, 64)
		return err == nil
	}
	n := len(tokens)
	for i := 0 ; i < n ; i++ {
		res := tokens[i]
		if isNum(res) {
			num,_ := strconv.Atoi(res)
			stack = append(stack,num)
		}else{
			ns := len(stack)
			res1,res2 := stack[ns-2],stack[ns-1]
			switch res {
			case "+":
				res1 = res1+res2
			case "-":
				res1 = res1-res2
			case "*":
				res1 = res1*res2
			case "/":
				res1 = res1/res2
			}
			stack[ns-2] = res1
			stack = stack[:ns-1]
		}
	}
	return stack[0]
}
// @lc code=end

