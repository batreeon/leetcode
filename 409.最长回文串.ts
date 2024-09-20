/*
 * @lc app=leetcode.cn id=409 lang=typescript
 *
 * [409] 最长回文串
 */

// @lc code=start
function longestPalindrome(s: string): number {
    let set = new Set<string>();
    let result: number = 0;
    for (let c of s) {
        if (!set.has(c)) {
            set.add(c);
        }else{
            set.delete(c);
            result += 2;
        }
    }
    if (set.size > 0) {
        result++;
    } 
    return result;
};
// @lc code=end

