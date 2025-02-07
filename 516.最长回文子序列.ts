/*
 * @lc app=leetcode.cn id=516 lang=typescript
 *
 * [516] 最长回文子序列
 */

// @lc code=start
// 字符串，动态规划
function longestPalindromeSubseq(s: string): number {
    let n = s.length
    let dp: number[][] = new Array(n).fill(0).map(() => new Array(n).fill(0));

    for (let l = 1; l <= n; l++) {
        for (let i = 0; i + l - 1 < n; i++) {
            let j = i+l-1;
            if (l === 1) {
                dp[i][j] = 1;
            }else if (l === 2 && s[i]===s[j]){
                dp[i][j] = 2;
            }else if (s[i] === s[j]) {
                    dp[i][j] = dp[i+1][j-1] + 2;
            }else{
                dp[i][j] = Math.max(dp[i+1][j], dp[i][j-1]);
            }
        }
    }
    return dp[0][n-1];
};
// @lc code=end

