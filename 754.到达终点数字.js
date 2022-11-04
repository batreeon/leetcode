/*
 * @lc app=leetcode.cn id=754 lang=javascript
 *
 * [754] 到达终点数字
 */

// @lc code=start
/**
 * @param {number} target
 * @return {number}
 */
/*
如果要走n步
那么就是 1 2 3 4 5 。。。 n-1 n
假设有x个正的，则还有（n-x）个负的
想不出来的
*/
var reachNumber = function(target) {
    // 抄的 https://leetcode.cn/problems/reach-a-number/solutions/1947419/by-ac_oier-o4ze/
    if (target < 0) {
        target = -target;
    }
    let k = Math.floor(Math.sqrt(2 * target));
    let s = k*(k+1)/2;

    while(s < target || (s-target)%2 == 1) {
        k++;
        s = k*(k+1)/2;
    }
    return k;
};
// @lc code=end

