/*
 * @lc app=leetcode.cn id=383 lang=typescript
 *
 * [383] 赎金信
 */

// @lc code=start
function canConstruct(ransomNote: string, magazine: string): boolean {
    let m = new Map();
    for (let i = 0; i < magazine.length; i++) {
        if (m.has(magazine[i])) {
            m.set(magazine[i], m.get(magazine[i]) + 1);
        } else {
            m.set(magazine[i], 1);
        }
    }
    for (let i = 0; i < ransomNote.length; i++) {
        if (m.has(ransomNote[i])) {
            if (m.get(ransomNote[i]) > 0) {
                m.set(ransomNote[i], m.get(ransomNote[i]) - 1);
                if (m.get(ransomNote[i]) === 0) {
                    m.delete(ransomNote[i]);
                }
            }
        } else {
            return false;
        }
    }
    return true;
};
// @lc code=end

