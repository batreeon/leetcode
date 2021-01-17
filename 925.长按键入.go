/*
 * @lc app=leetcode.cn id=925 lang=golang
 *
 * [925] 长按键入
 */

// @lc code=start
func isLongPressedName(name string, typed string) bool {
	if len(name) == 0 {
		return true
	}
	if name[0] != typed[0] || name[len(name)-1] != typed[len(typed)-1] {
		return false
	}
	var nowletter uint8
	ni,ti := 0,0
	for ;ni < len(name);ni,ti = ni+1,ti+1 {
		if ti >= len(typed) {
			return false
		}
		nowletter = name[ni]
		if(typed[ti] != nowletter) {
			return false
		}
		for ;ni < len(name)-1 && name[ni+1] == nowletter ;ni++{
			if typed[ti+1] != nowletter {
				return false
			}
			ti++
		}
		for ;ti < len(typed)-1 && typed[ti+1] == nowletter ;ti++{
		}
	}
	return true
}
// @lc code=end

