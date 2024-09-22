/*
 * @lc app=leetcode.cn id=164 lang=typescript
 *
 * [164] 最大间距
 */

// @lc code=start
function maximumGap(nums: number[]): number {
    // sort默认是按字符序进行排序
    nums = nums.sort(function(a, b){return a - b});
    let maxGap: number = 0;
    for (let i = 1; i < nums.length; i++) {
        let gap: number =  nums[i] - nums[i-1]
        if (gap > maxGap) {
            maxGap = gap;
        }
    }
    return maxGap;
};
// @lc code=end

