/*
 * @lc app=leetcode.cn id=242 lang=javascript
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function(s, t) {
    let chs = new Array();
    for (let i = 0; i < 26; i++) {
        chs[i] = 0;
    }
    let aidx = 'a'.charCodeAt();
    for (let ss of s) {
        chs[ss.charCodeAt() - aidx]++;
    }
    for (let tt of t) {
        chs[tt.charCodeAt() - aidx]--;
    }
    for (let ch of chs) {
        if (ch != 0) {
            return false;
        }
    }
    return true;
};
// @lc code=end

