/*
 * @lc app=leetcode.cn id=198 lang=typescript
 *
 * [198] 打家劫舍
 */

// @lc code=start
// 动态规划，好题
function rob(nums: number[]): number {
    let dp1 = 0;
    let dp2 = 0;
    nums.forEach((num) => {
        // max中的两个值，分别代表不偷第i家 / 偷第i家，偷第i家则不能偷第i-1家
        let dpi = Math.max(dp2, dp1+num);
        dp1 = dp2;
        dp2 = dpi;
    })
    return dp2;
};
// @lc code=end

