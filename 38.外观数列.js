/*
 * @lc app=leetcode.cn id=38 lang=javascript
 *
 * [38] 外观数列
 */

// @lc code=start
/**
 * @param {number} n
 * @return {string}
 */
let ss = [];
var countAndSay = function(n) {
    if (n === 1) {
        ss.push('1');
    }
    if (ss.length === n) {
        return ss[n - 1];
    } 
    for (;ss.length < n;) {
        ss.push(describe(ss[ss.length - 1]));
    }
    return ss[n - 1]
};

let describe = function(s) {
    let result = '';
    let i = 0;
    for (let j = i; i < s.length; ) {
        for (; j < s.length && s[j] === s[i]; ) {
            j++;
        }
        result += (j-i).toString();
        result += s[i];
        i = j;
    } 
    return result;
};
// @lc code=end

