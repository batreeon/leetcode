/*
 * @lc app=leetcode.cn id=908 lang=typescript
 *
 * [908] 最小差值 I
 */

// @lc code=start
function smallestRangeI(nums: number[], k: number): number {
    nums = nums.sort((a,b) => {return a - b});
    let min:number, max:number;
    let l:number = nums.length;
    min = nums[0];
    max = nums[l-1];
    if ((max - min) <= 2*k) {
        return 0;
    }
    return max - min - 2*k;
};
// @lc code=end

