/*
 * @lc app=leetcode.cn id=468 lang=golang
 *
 * [468] 验证IP地址
 */

// @lc code=start
func validIPAddress(queryIP string) string {
	if isIPV4(queryIP) {
		return string("IPv4")
	}
	if isIPV6(queryIP) {
		return string("IPv6")
	}
	return string("Neither")
}

func isIPV4(ip string) bool {
	if !strings.Contains(ip, ".") || strings.Contains(ip, ":") {
		return false
	}
	nums4 := strings.Split(ip, ".")
	if len(nums4) != 4 {
		return false
	}
	for _, num4 := range nums4 {
		if !isNum4(num4) {
			return false
		}
	}
	return true
}

func isIPV6(ip string) bool {
	if !strings.Contains(ip, ":") || strings.Contains(ip, ".") {
		return false
	}
	nums6 := strings.Split(ip, ":")
	if len(nums6) != 8 {
		return false
	}
	for _, num6 := range nums6 {
		if !isNum6(num6) {
			return false
		}
	}
	return true
}

func isNum4(num string) bool {
	if strings.Contains(num, "e") {
		return false
	}
	if len(num) > 3 || len(num) == 0 {
		return false
	}
	if num[0] == '0' && len(num) != 1 {
		return false
	}
	n, err := strconv.Atoi(num)
	if err != nil {
		return false
	} 
	if n < 0 || n > 255 {
		return false
	}
	return true
}

func isNum6(num string) bool {
	if len(num) > 4 || len(num) == 0 {
		return false
	}
	for _, n := range num {
		if !((n >= '0' && n <= '9') || (n >= 'a' && n <= 'f') || (n >= 'A' && n <= 'F')) {
			return false
		}
	}
	return true 
}
// @lc code=end

