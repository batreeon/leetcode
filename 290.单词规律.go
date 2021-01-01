/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	// mp := map[string][]int{}
	// for i,a := range pattern {
	// 	mp[string(a)] = append(mp[string(a)],i)
	// }
	// mmp := map[int][]int{}
	// for _,v := range mp {
	// 	if len(v) > 1{
	// 		mmp[v[0]] = v[1:]
	// 	}
	// }
	// sstr := strings.Split(s," ")
	// mp1 := map[string][]int{}
	// for i,a := range sstr {
	// 	mp1[string(a)] = append(mp1[string(a)],i)
	// }
	// mmp1 := map[int][]int{}
	// for _,v := range mp {
	// 	if len(v) > 1{
	// 		mmp1[v[0]] = v[1:]
	// 	}
	// }
	// for k,v := range mmp {
	// 	for i,j := range v {
	// 		if j != mmp1[k][i] {
	// 			return false
	// 		}
	// 	}
	// }
	// return true
	p2s := map[byte]string{}
	//如果只用一个map只能判断单项的，如"abba"
	//"dog dog dog dog"
	//判断双向需要两个数组
	s2p := map[string]byte{}
	ss := strings.Split(s," ")
	if len(pattern) != len(ss) {
		return false
	}
	for i,word := range ss {
		p := pattern[i]
		// if _,ok1 := p2s[p];!ok {
		// 	p2s[p] = word
		// }else if p2s[p] != word {
		// 	return false
		// }
		if p2s[p] != "" && p2s[p] != word || s2p[word] > 0 && s2p[word] != p {
			return false
		}//要么两边都是首次出现。若出现了但不配对则错误。只有这两种情况
		p2s[p] = word
		s2p[word] = p
	}
	return true
}
// @lc code=end

