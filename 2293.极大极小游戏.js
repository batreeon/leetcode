/*
 * @lc app=leetcode.cn id=2293 lang=javascript
 *
 * [2293] 极大极小游戏
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number}
 */
var minMaxGame = function(nums) {
    return f(nums)[0]
};

function f(nums) {
    if (nums.length == 1) {
        return nums
    }
    let newNums = new Array();
    for (let i = 0; i < nums.length/2; i++) {
        newNums[i] = (i % 2 == 0) ? Math.min(nums[i*2], nums[i*2+1]) : Math.max(nums[i*2], nums[i*2+1])
    }
    return f(newNums)
}
// @lc code=end

