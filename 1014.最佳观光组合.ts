/*
 * @lc app=leetcode.cn id=1014 lang=typescript
 *
 * [1014] 最佳观光组合
 */

// @lc code=start
// 看的答案，维护j左侧最大的 vi + i
function maxScoreSightseeingPair(values: number[]): number {
    let maxviaddi = 0;
    let result = 0;
    for (let j = 0; j < values.length; j++) {
        result = Math.max(result, maxviaddi + values[j] - j);
        maxviaddi = Math.max(maxviaddi, values[j] + j);
    }
    return result;
};
// @lc code=end

