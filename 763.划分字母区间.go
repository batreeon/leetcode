/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */

// @lc code=start
func partitionLabels(S string) []int {
	newletter := make(map[uint8]bool)
	var ans []int
	lastindex := strings.LastIndex(S,string(S[0]))
	for i := 0; (lastindex <= len(S)-1) && (i <= len(S)-1); i++ {
		lastindex := strings.LastIndex(S,string(S[i]))
		firstindex := i
		//lastindex := strings.LastIndex(S,S[i])
		for ;i < lastindex; i++ {
			if !newletter[S[i]] {
				newletter[S[i]] = true
				if index := strings.LastIndex(S,string(S[i])); index > lastindex {
					lastindex = index
				}
			}			
		}
		ans =  append(ans,int(lastindex-firstindex+1))
	}
	return ans
}
// @lc code=end

