/*
 * @lc app=leetcode.cn id=3180 lang=typescript
 *
 * [3180] 执行操作可获得的最大总奖励 I
 */

// @lc code=start
function maxTotalReward(rewardValues: number[]): number {
    rewardValues.sort((x,y) => {return x - y});
    let result: number = 0;
    let dp: number[] = new Array(rewardValues.length);
    for (let i = 0; i < rewardValues.length; i++) {
        
    } 
};
// @lc code=end

