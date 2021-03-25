/*
 * @lc app=leetcode.cn id=393 lang=golang
 *
 * [393] UTF-8 编码验证
 */

// @lc code=start
package main
func validUtf8(data []int) bool {
	var validutf8 func(data []int) bool
	validutf8 = func(data []int) bool {
		if len(data) == 0 {
			return true
		}
		// 不采用注释中的判断条件，因为可能对0进行移位，会报错
		// if (data[0]>>7)&1 == 0 {
		// 判断最高位是否为0，若为0，则为1字节
		if data[0]&(1<<7) ==0 {
			return validutf8(data[1:])
		}

		// 最高位不为0，检查高位共有多少个1
		n := 6
		// for data[0]>>n != 0 {
		for data[0]&(1<<n) != 0 {
			// 最多只能有4字节，若超过了则不合法
			if n == 3 {
				return false
			}
			n--
		}

		// 若n==6,说明次高位为0，即只有一位1
		// n记录字节数
		n = 7-n
		// 1字节对应的0开头，若存在10为首字节开头，则不合法
		if n == 1 {
			return false
		}

		// 判断字节数是否够，并保证后续遍历时下标合法
		if len(data) < n {
			return false
		}

		// 检查非首字节
		for i := 1 ; i < n ; i++ {
			// if (data[i]>>7)&1 != 1 || (data[i]>>6)&1 != 0 {
			// 若不为10开头，则报错
			if data[i]&(1<<7) != (1<<7) || data[i]&(1<<6) != 0 {
				return false
			}
		}
		// 前面的都合法，递归检查剩余数组
		return validutf8(data[n:])
	}
	return validutf8(data)
}
// @lc code=end

