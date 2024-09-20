/*
 * @lc app=leetcode.cn id=387 lang=typescript
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
function firstUniqChar(s: string): number {
    let map = new Map();
    for (let i = 0; i < s.length; i++) {
        let nums = 0;
        if (map.has(s[i])) {
            nums = map.get(s[i]);
        }
        map.set(s[i], nums + 1);
    }

    for (let i = 0; i < s.length; i++) {
        if (map.get(s[i]) === 1) {
            return i;
        }
    }
    return -1;
};
// @lc code=end

