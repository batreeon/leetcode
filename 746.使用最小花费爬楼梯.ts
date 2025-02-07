/*
 * @lc app=leetcode.cn id=746 lang=typescript
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
// 初级动态规划，撒撒水
function minCostClimbingStairs(cost: number[]): number {
    let dp0 = cost[0];
    let dp1 = cost[1];
    
    for (let i = 2; i <= cost.length; i++) {
        if (i === cost.length) {
            return Math.min(dp0, dp1)
        }
        let dpi = Math.min(dp0, dp1) + cost[i];
        dp0 = dp1;
        dp1 = dpi;
    }
};
// @lc code=end

