/*
 * @lc app=leetcode.cn id=1310 lang=golang
 *
 * [1310] 子数组异或查询
 */

// @lc code=start
func xorQueries(arr []int, queries [][]int) []int {
	xors := make([]int,len(arr)+1)
	for i := 0 ; i < len(xors)-1 ; i++ {
		xors[i+1] = xors[i] ^ arr[i]
	}
	result := make([]int,len(queries))
	for i := 0 ; i < len(queries) ; i++ {
		result[i] = xors[queries[i][0]] ^ xors[queries[i][1]+1]
		// arr[0] ^ ... arr[left-1] ^ arr[0] ^ ... arr[right]
		// = arr[left] ^ ... ^ arr[right]
	}
	return result
}
// @lc code=end

