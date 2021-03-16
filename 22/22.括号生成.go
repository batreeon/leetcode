/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
package main
func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}

	ans := []string{}

	var gp func(res []byte,lR,rR,status int)
	gp = func(res []byte , lR,rR,status int) {
		if lR == 0 && rR == 0 {
			ans = append(ans,string(res))
			return
		}
		if status == 0 {
			res = append(res,'(')
			gp(res,lR-1,rR,status+1)
		}else{
			if lR > 0 {
				res = append(res,'(')
				gp(res,lR-1,rR,status+1)
				res = res[:len(res)-1]
			}
			res = append(res,')')
			gp(res,lR,rR-1,status-1)
		}
	}
	gp([]byte{'('},n-1,n,1)
	return ans
}
// @lc code=end