/*
 * @lc app=leetcode.cn id=1662 lang=javascript
 *
 * [1662] 检查两个字符串数组是否相等
 */

// @lc code=start
/**
 * @param {string[]} word1
 * @param {string[]} word2
 * @return {boolean}
 */
var arrayStringsAreEqual = function(word1, word2) {
    let s1 = "", s2 = "";
    for (w1 of word1) {
        s1 += w1;
    }
    for (w2 of word2) {
        s2 += w2;
    }
    if (s1.length != s2.length) {
        return false;
    }
    return compareString(s1, s2);
};

function compareString(s1, s2) {
    for (let i = 0; i < s1.length; i++) {
        if (s1[i] != s2[i]) {
            return false
        }
    }
    return true
}
// @lc code=end

