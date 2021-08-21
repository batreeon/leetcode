/*
 * @lc app=leetcode.cn id=443 lang=golang
 *
 * [443] 压缩字符串
 */

// @lc code=start
package main

import (
	"strconv"
)

func compress(chars []byte) int {
	for i := 0 ; i < len(chars) ; {
		j := i+1
		for ; j < len(chars) ; j++ {
			if chars[j] != chars[i] {
				break
			}
		}
		// chars[j]是第一个与chars[i]不同的元素，那么chars[i]的重复次数就是j-i
		count := j-i
		if count == 1 {
			i = j
			continue
		}else{
			// count不为1，把count转为字符串
			countb := strconv.Itoa(count)
			// chars[i]保留，后面跟上重复次数字符串countb
			copy(chars[i+1:],[]byte(countb))
			// 好了，接着把多余的chars[i]去掉，只需要保留首字符和重复次数就够了
			// 只需要保留chars[i]和后面len(countb)长度的内容就够了 
			// 将chars[j:]后面的内容往前移，覆盖掉不需要的
			copy(chars[i+1+len(countb):],chars[j:])
			// 往前移了，那么末尾就有一些多余的不需要的了 
			// 由上式得(j-i-1-len(countb))是往前移动的单位数，也就是chars末尾多余的长度
			chars = chars[:len(chars)-(j-i-1-len(countb))]
			// 下一个考虑chars[j]了，但原来的chars[j]已经被我们移动到了chars[i+1+len(countb)]处
			i = i + 1 + len(countb)
		}
	}
	return len(chars)
}
// @lc code=end

