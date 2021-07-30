/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel表列序号
 */

// @lc code=start
func titleToNumber(columnTitle string) int {
    var res int
    for i := range columnTitle {
        res = res*26 + int(columnTitle[i]-'A')+1
    }
    return res
}
// @lc code=end

