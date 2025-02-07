/*
 * @lc app=leetcode.cn id=70 lang=typescript
 *
 * [70] 爬楼梯
 */

// @lc code=start
// 最简单的动态规划
function climbStairs(n: number): number {
    if (n <= 2) {
        return n;
    }
    let dp1 = 1;
    let dp2 = 2;
    for (let i = 3; i <= n; i++) {
        let dpi = dp1 + dp2;
        dp1 = dp2;
        dp2 = dpi;
    }
    return dp2;
};
// @lc code=end

