/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
// 解1
// func groupAnagrams(strs []string) [][]string {
// 	//把string按字母顺序排序后作为键
// 	mp := map[string][]string{} //k是string,v是string数组
// 	for _,str := range strs {
// 		s := []byte(str)
// 		sort.Slice(s, func(i,j int) bool {return s[i] < s[j]})
// 		sortedStr := string(s) //sortedStr就是将原string排序后的结果
// 		mp[sortedStr] = append(mp[sortedStr],str)
// 	}
// 	ans := make([][]string,0,len(mp))
// 	for _,v := range mp {
// 		ans = append(ans,v)
// 	}
// 	return ans
// }

// 解2
func groupAnagrams(strs []string) [][]string {
	//对每个单词中的每个字母计数，用一个26元素的int数组来记录，作为map的键
	mp := map[[26]int][]string{} //k是[26]int,v是string数组
	for _,str := range strs {
		cnt := [26]int{}
		for _,c := range str {
			cnt[c-'a']++
		}
		mp[cnt] = append(mp[cnt],str)
	}
	ans := make([][]string,0,len(mp))
	for _,v := range mp {
		ans = append(ans,v)
	}
	return ans
}
// @lc code=end

