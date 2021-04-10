/*
 * @lc app=leetcode.cn id=990 lang=golang
 *
 * [990] 等式方程的可满足性
 */

// @lc code=start
func equationsPossible(equations []string) bool {
	// 先记录每个字母相等的字母序列
	// 再记录每个字母不相等的字母序列
	// 用color记录每个字母的标色，初始为未标色0
	equal := map[byte][]byte{}
	inequal := map[byte][]byte{}
	color := map[byte]int{}
	for _,eq := range equations {
		color[eq[0]] = 0
		color[eq[3]] = 0
		if eq[1] == '=' && eq[0] != eq[3] {
			equal[eq[0]] = append(equal[eq[0]],eq[3])
			equal[eq[3]] = append(equal[eq[3]],eq[0])
		}else if eq[1] == '!' {
			if eq[0] == eq[3] {
				return false
			}
			inequal[eq[0]] = append(inequal[eq[0]],eq[3])
			inequal[eq[3]] = append(inequal[eq[3]],eq[0])
		}
	}

	// 利用dfs将所有相等的，标上相同的颜色，每一次dfs把同一种颜色的标记完
	// 然后colorNum++，再标下一种颜色
	var dfs func(v byte , colorNnm int)
	dfs = func(v byte,colorNum int) {
		if color[v] != 0 {
			return
		}
		color[v] = colorNum
		for _,next := range equal[v] {
			dfs(next,colorNum)
		}
	}
	colorNum := 0
	for k := range equal {
		if color[k] == 0 {
			colorNum++
			dfs(k,colorNum)
		}
	}
	
	// 然后利用不相等关系，检查有没有具有不相等关系，
	// 却还标了相同颜色的，这会产生矛盾，直接返回false
	// 若其中一个未标色，那就是说该字母没有出现在任何一组相等关系中，
	// 那么他的所有不等关系都是成立，跳过检查下一个
	for k,inequals := range inequal {
		nowColor := color[k]
		if nowColor == 0 {
			continue
		}
		for _,next := range inequals {
			if color[next] == 0 {
				continue
			}
			if color[next] == nowColor {
				return false
			}
		}
	}

	return true
}
// @lc code=end

