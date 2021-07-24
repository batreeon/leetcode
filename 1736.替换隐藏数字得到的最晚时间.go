/*
 * @lc app=leetcode.cn id=1736 lang=golang
 *
 * [1736] 替换隐藏数字得到的最晚时间
 */

// @lc code=start
func maximumTime(time string) string {
	bytes := []byte(time)
	if bytes[4] == '?' {
		bytes[4] = '9'
	}
	if bytes[3] == '?' {
		bytes[3] = '5'
	}
	if bytes[1] == '?' && bytes[0] == '?' {
		bytes[1],bytes[0] = '3','2'
	}else if bytes[1] == '?' {
		if bytes[0] == '2' {
			bytes[1] = '3'
		}else{
			bytes[1] = '9'
		}
	}else if bytes[0] == '?' {
		if bytes[1] - '0' <= 3 {
			bytes[0] = '2'
		}else{
			bytes[0] = '1'
		}
	}

	return string(bytes)
}
// @lc code=end

