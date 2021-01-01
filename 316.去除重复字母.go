/*
 * @lc app=leetcode.cn id=316 lang=golang
 *
 * [316] 去除重复字母
 */

// @lc code=start
func removeDuplicateLetters(s string) string {
	// mp := map[byte]int{}
	// //记录留下的字符的坐标
	// bs := []byte(s)
	// for i,l := range bs {
	// 	if first,ok := mp[l]; !ok || bs[first+1] < l {
	// 		mp[l] = i;
	// 	}
	// }
	// no := []int{}
	// for _,v := range mp {
	// 	no = append(no,v)
	// }
	// sort.Ints(no)
	// ans := []byte{}
	// for _,v := range no{
	// 	ans = append(ans,bs[v])
	// }
	// return string(ans)
	// 错误解法，例子　bcabc就通不过

    left := [26]int{}
    for _, ch := range s {
        left[ch-'a']++
    }
    stack := []byte{}
    inStack := [26]bool{}
    for i := range s {
        ch := s[i]
        if !inStack[ch-'a'] {
            for len(stack) > 0 && ch < stack[len(stack)-1] {
                last := stack[len(stack)-1] - 'a'
                if left[last] == 0 {
                    break
                }
                stack = stack[:len(stack)-1]
                inStack[last] = false
            }
            stack = append(stack, ch)
            inStack[ch-'a'] = true
        }
        left[ch-'a']--
    }
    return string(stack)
}
// @lc code=end

