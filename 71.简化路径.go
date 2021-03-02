/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
func simplifyPath(path string) string {
	array := strings.Split(path,"/")
	if len(array) == 0 {
		return string("/")
	}
	if array[0] == string("..") {
		return string("/")
	}
	for i := 0 ; i < len(array) ; i++ {
		if array[i] != string(".") {
			array = array[i+1:]
			break
		}
	}
	now := []string{}
	for len(array) > 0 {
		if len(now) == 0 && array[0] == string("..") {
			return string("/")
		}
		if array[0] == string(".") || array[0] == string("") {
		}else if array[0] == string(".."){
			now = now[1:]
		}else{
			now = append(now,array[0])
		}
		array = array[1:]
	}
	ans := strings.Join([]string{"/",strings.Join(now,"/")},"")
	return ans
}
// @lc code=end

