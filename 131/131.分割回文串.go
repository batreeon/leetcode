/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
package main

func partition(s string) [][]string {
	// 回溯法
	/*
	n := len(s)
	f := make([][]bool,n)
	fRes := make([]bool,n)
	for i := 0 ; i < n ; i++ {
		fRes[i] = true
	}
	for i := 0 ; i < n ; i++ {
		// 不可以直接写f[i] = fRes，那样的话每一行都是一样的
		f[i] = make([]bool,n)	//这一句不能省
		// 不能直接这样写copy(f[i],fRes)
		copy(f[i],fRes)

	}
	// f[i][j]标记s[i:j+1]是否是回文，若i > j则是回文，否则f[i][j] = s[i]==s[j] && f[i+1][j-1]
	for i := n-1 ; i >= 0 ; i-- {
		for j := i + 1 ; j < n ; j++ {
			f[i][j] = ((s[i] == s[j]) && f[i+1][j-1])
		}
	}

	result := [][]string{}
	res := []string{}
	var dfs func(begin int)
	dfs = func(begin int) {
		if begin == n {
			resCopy := make([]string,len(res))
			copy(resCopy,res)
			result = append(result,resCopy)
		}else{
			for i := begin ; i < n ; i++ {
				if f[begin][i] {
					res = append(res,s[begin:i+1])
					dfs(i+1)
					res = res[:len(res)-1]
				}
			}
		}
	}
	dfs(0)
	return result
	*/

	// 回溯，记忆化
	n := len(s)
	f := make([][]bool,n)
	fRes := make([]bool,n)
	for i := 0 ; i < n ; i++ {
		fRes[i] = false
	}
	for i := 0 ; i < n ; i++ {
		// 不可以直接写f[i] = fRes，那样的话每一行都是一样的
		f[i] = make([]bool,n)	//这一句不能省
		// 不能直接这样写copy(f[i],fRes)
		copy(f[i],fRes)

	}
	// f[i][j]标记s[i:j+1]是否是回文，若i > j则是回文，否则f[i][j] = s[i]==s[j] && f[i+1][j-1]
	var isPalindrome func(i,j int) bool
	isPalindrome = func(i,j int) bool {
		if f[i][j] {
			return true
		}
		if i >= j {
			f[i][j] = true
			return true
		}
		f[i][j] = ((s[i] == s[j]) && isPalindrome(i+1,j-1))
		return f[i][j]
	}

	result := [][]string{}
	res := []string{}
	var dfs func(begin int)
	dfs = func(begin int) {
		if begin == n {
			resCopy := make([]string,len(res))
			copy(resCopy,res)
			result = append(result,resCopy)
		}else{
			for i := begin ; i < n ; i++ {
				if isPalindrome(begin,i) {
					res = append(res,s[begin:i+1])
					dfs(i+1)
					res = res[:len(res)-1]
				}
			}
		}
	}
	dfs(0)
	return result
}
// @lc code=end

