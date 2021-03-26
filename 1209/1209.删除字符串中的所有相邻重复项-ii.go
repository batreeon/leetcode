/*
 * @lc app=leetcode.cn id=1209 lang=golang
 *
 * [1209] 删除字符串中的所有相邻重复项 II
 */

// @lc code=start
package main
func removeDuplicates(s string, k int) string {
	sByte := []byte(s)
	l := len(s)
	begin:= 0

	// 初始元素，相同元素数为1
	n := 1
	// 后半部是不变的，因此若可以遍历到尾，就结束了
	for i := begin+1 ; i < l ; i++ {
		// 相同序列长度计数
		if sByte[i] == sByte[i-1] {
			n++
			// 符合删除条件
			if n == k {
				// 用前面的元素将删除掉的元素覆盖掉
				copy(sByte[begin+k:i+1],sByte[begin:i-k+1])
				// 新的头下标
				begin = begin+k
				// 从新的头开始遍历
				i = begin
				n = 1
			}
		}else{
			// 新的序列
			n = 1
		}
	}
	return string(sByte[begin:])
}
// @lc code=end

