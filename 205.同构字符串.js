/*
 * @lc app=leetcode.cn id=205 lang=javascript
 *
 * [205] 同构字符串
 */

// @lc code=start
/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isIsomorphic = function(s, t) {
    let s2t = new Map();
    let t2s = new Set();

    for (let i = 0; i < s.length; i++) {
        // 检查是否有多对一
        if (!s2t.has(s.charAt(i))) {
            if (t2s.has(t.charAt(i))) {
                return false;
            }
            s2t.set(s.charAt(i), t.charAt(i));
            t2s.add(t.charAt(i));
        // 检查是否有一对多
        }else if (t.charAt(i) != s2t.get(s.charAt(i))) {
            return false
        }
    }
    return true;
};
// @lc code=end

