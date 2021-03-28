package main
/*
给你一个字符串 s ，它包含一些括号对，每个括号中包含一个 非空 的键。

比方说，字符串 "(name)is(age)yearsold" 中，有 两个 括号对，分别包含键 "name" 和 "age" 。
你知道许多键对应的值，这些关系由二维字符串数组 knowledge 表示，其中 knowledge[i] = [keyi, valuei] ，表示键 keyi 对应的值为 valuei 。

你需要替换 所有 的括号对。当你替换一个括号对，且它包含的键为 keyi 时，你需要：

将 keyi 和括号用对应的值 valuei 替换。
如果从 knowledge 中无法得知某个键对应的值，你需要将 keyi 和括号用问号 "?" 替换（不需要引号）。
knowledge 中每个键最多只会出现一次。s 中不会有嵌套的括号。

请你返回替换 所有 括号对后的结果字符串。
*/
// import "strings"
func evaluate(s string, knowledge [][]string) string {
	/*
	// 替换居然超时了
	m := map[string]string{}
	for _, k := range knowledge {
		m[k[0]] = k[1]
	}
	stack := []int{}
	*/

	// for i, ss := range s {
	// 	if ss == '(' || ss == ')' {
	// 		stack = append(stack, i)
	// 	}
	// }

	/*
	mm := map[string]string{}
	for i, ss := range s {
		if ss == '(' {
			stack = append(stack, i)
		}
		if ss == ')' {
			left,right := stack[0],i
			if _,ok := mm[s[left:right+1]] ; !ok {
				if v,ok := m[s[left+1:right]] ; ok {
					mm[s[left:right+1]] = v
				}else{
					mm[s[left:right+1]] = string("?")
				}
			}
			stack = stack[1:]
		}
	}
	for k,v := range mm {
		s = strings.ReplaceAll(s,k,v)
	}
	return s
	*/

	// sCopy := string(s)
	// for len(stack) > 0 {
	// 	left, right := stack[0], stack[1]
	// 	if v, ok := m[s[left+1:right]]; ok {
	// 		sCopy = strings.ReplaceAll(sCopy,s[left:right+1],v)
	// 	}else{
	// 		sCopy = strings.ReplaceAll(sCopy,s[left:right+1],"?")
	// 	}
    //     stack = stack[2:]
	// }
	// return sCopy

	// 拼接
	m := map[string]string{}
	for _, k := range knowledge {
		m[k[0]] = k[1]
	}
	var ans string
	for i := 0 ; i < len(s) ; i++ {
		if s[i] == '(' {
			var key string
			for j := i+1 ; j < len(s) ; j++ {
				if s[j] != ')' {
					key += string(s[j])
				}else{
					i = j
					break
				}
			}
			if v,ok := m[key] ; ok {
				ans += string(v)
			}else{
				ans += string("?")
			}
		}else{
			ans += string(s[i])
		}
	}
	return ans
}