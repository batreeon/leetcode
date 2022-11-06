/*
 * @lc app=leetcode.cn id=268 lang=javascript
 *
 * [268] 丢失的数字
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number}
 */
var missingNumber = function(nums) {
    let len = nums.length;
    let sum = len * (len+1) / 2;

    for (let num of nums) {
        sum -= num;
    }
    return sum;
};
// @lc code=end

