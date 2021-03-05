/*
 * @lc app=leetcode.cn id=636 lang=golang
 *
 * [636] 函数的独占时间
 */

// @lc code=start
package main

import (
	"strconv"
	"strings"
)


func exclusiveTime(n int, logs []string) []int {

	//看的官方解答，太难了
	
	//记录id,栈长
	stack := []int{}
	sl := 0

	ans := make([]int,n)

	begin := strings.Index(logs[0],":")
	end := strings.LastIndex(logs[0],":")
	id,_ := strconv.Atoi(logs[0][:begin])
	stack = append(stack,id)
	sl++

	//记录栈顶元素开始时间
	prev,_ := strconv.Atoi(logs[0][end+1:])

	l := len(logs)
	for i := 1 ; i < l ; i++ {
		begin = strings.Index(logs[i],":")
		end = strings.LastIndex(logs[i],":")
		signal := logs[i][begin+1:end]
		time,_ := strconv.Atoi(logs[i][end+1:])
		id,_ = strconv.Atoi(logs[i][:begin])
		if signal == "start" {
			if sl > 0 {
				ans[stack[sl-1]] += time - prev
			}
			stack = append(stack,id)
			sl++
			prev = time
		}else{
			ans[stack[sl-1]] += time - prev + 1
			stack = stack[:sl-1]
			sl--
			prev = time + 1
		}
	}
	return ans
}
// @lc code=end

