/*
 * @lc app=leetcode.cn id=1415 lang=golang
 *
 * [1415] 长度为 n 的开心字符串中字典序第 k 小的字符串
 */

// @lc code=start
package main
func getHappyString(n int, k int) string {
	ans := []string{}
	s := string("abc")
    var gethappystring func(res []byte)
	gethappystring = func(res []byte) {
		if len(res) == n {
			resCopy := make([]byte,n)
			copy(resCopy,res)
			ans = append(ans,string(resCopy))
			return
		}
		for i := 0 ; i < 3 ; i++ {
			if (len(res) > 0 && s[i] != res[len(res)-1]) || len(res) == 0 {
				res = append(res,s[i])
				gethappystring(res)
				res = res[:len(res)-1]
			}
		}
	}
	gethappystring([]byte{})
	if len(ans) < k {
		return string("")
	}
	return ans[k-1]
}
// @lc code=end

