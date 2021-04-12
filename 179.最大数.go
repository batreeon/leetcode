/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 */

// @lc code=start
type sortstrings []string
func (s sortstrings) Len() int {return len(s)}
func (s sortstrings) Less(i,j int) bool {return strings.Join([]string{s[i],s[j]},"") > strings.Join([]string{s[j],s[i]},"") }
func (s sortstrings) Swap(i,j int) {s[i],s[j] = s[j],s[i]}

func largestNumber(nums []int) string {
	n := len(nums)
	strnums := make([]string,n)
	for i := 0 ; i < n ; i++ {
		strnums[i] = strconv.Itoa(nums[i])
	}
	sort.Sort(sortstrings(strnums))
	result := strings.Join(strnums,"")
	if resultNum,_ := strconv.Atoi(result) ; resultNum == 0 {
		result = string("0")
	}
	return result
}
// @lc code=end

