/*
 * @lc app=leetcode.cn id=219 lang=javascript
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 */
var containsNearbyDuplicate = function(nums, k) {
    let idxs = new Map();
    for (let i = 0; i < nums.length; i++) {
        if (idxs.has(nums[i])) {
            if (i - idxs.get(nums[i])<= k) {
                return true;
            }
        }
        idxs.set(nums[i], i);
    }
    return false;
};
// @lc code=end

