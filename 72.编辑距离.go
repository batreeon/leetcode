/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	n1,n2 := len(word1),len(word2)
	f1 := make([]int,n2+1)
	for i := 0 ; i < n2+1 ; i++ {
		f1[i] = i
	}
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	for i := 1 ; i < n1+1 ; i++ {
		f := make([]int,n2+1)
		f[0] = i
		for j := 1 ; j < n2+1 ; j++ {
			if word1[i-1] == word2[j-1] {
				f[j] = min(f1[j-1],min(f[j-1]+1,f1[j]+1))
			}else{
				f[j] = min(f1[j-1]+1,min(f[j-1]+1,f1[j]+1))
			}
		}
		copy(f1,f)
	}
	return f1[n2]
}
// @lc code=end

