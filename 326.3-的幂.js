/*
 * @lc app=leetcode.cn id=326 lang=javascript
 *
 * [326] 3 的幂
 */

// @lc code=start
/**
 * @param {number} n
 * @return {boolean}
 */
// 神中神，抄的https://leetcode.cn/problems/power-of-three/solutions/1012126/gong-shui-san-xie-yi-ti-san-jie-shu-xue-8oiip/
var isPowerOfThree = function(n) {
    return n > 0 && 1162261467 % n == 0;
};
// @lc code=end

