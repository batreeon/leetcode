/*
 * @lc app=leetcode.cn id=1668 lang=javascript
 *
 * [1668] 最大重复子字符串
 */

// @lc code=start
/**
 * @param {string} sequence
 * @param {string} word
 * @return {number}
 */
var maxRepeating = function(sequence, word) {
    let result = 0;
    for (let i = 0; i < sequence.length-word.length+1; i++) {
        let k = 0;
        for (; sequence.slice(i,i+word.length) == word; i += word.length) {
            k++;
        }

        // 如果不判断k > 0；那就有可能一直原地踏步了
        // 每次前进了word.length，按如果某一次没有匹配上word，那么就要会滚一下
        // 将前进word.length改为前进前进一步（最外层循环）
        if (k > 0) {
            i -= word.length;
        }

        if (k > result) {
            result = k;
        }
    }
    return result
};
// @lc code=end

